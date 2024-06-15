package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/inview-team/sadko_indexer/internal/entities"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	uri     string
	mx      sync.RWMutex
}

type Config struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func New(config Config) (*Client, error) {
	c := new(Client)
	c.uri = fmt.Sprintf("amqp://%s:%s@%s:%d", config.User, config.Password, config.IP, config.Port)
	err := c.connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Rabbitmq")
	}
	return c, nil
}

func (c *Client) connect() error {
	var err error

	c.conn, err = amqp.Dial(c.uri)
	if err != nil {
		return err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return err
	}

	c.queue, err = c.channel.QueueDeclare(
		"tasks",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) reconnect() error {
	if err := c.connect(); err != nil {
		return err
	}
	return nil
}

func (c *Client) checkConnection() error {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.conn == nil || c.conn.IsClosed() {
		return c.reconnect()
	}
	return nil
}

func (c *Client) TagVideo(ctx context.Context, v *entities.Video) error {
	msg := NewTagMessage(v)
	data, err := json.Marshal(msg)

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	if err = c.checkConnection(); err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	tCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = c.channel.PublishWithContext(tCtx, "", c.queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
