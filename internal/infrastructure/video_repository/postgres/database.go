package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/inview-team/sadko_indexer/internal/entities"
)

type Client struct {
	client *pgxpool.Pool
}

type Config struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func New(ctx context.Context, config Config) (*Client, error) {
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.IP, config.Port, config.Database))
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return &Client{
		client: pool,
	}, nil
}

func (c *Client) NextID(_ context.Context) entities.VideoID {
	return entities.VideoID(uuid.New().String())
}

func (c *Client) Create(ctx context.Context, video *entities.Video) error {
	query := `INSERT INTO videos (id, url, description, related_vectors) VALUES ($1, $2, $3, $4)`

	_, err := c.client.Exec(ctx, query, video.ID, video.URL, video.Description, video.RelatedVectorIDs)
	if err != nil {
		return fmt.Errorf("failed to insert video: %v", err)
	}
	return nil
}

func (c *Client) GetByID(ctx context.Context, videoID string) (*entities.Video, error) {
	query := `SELECT id, url, description, vectors FROM videos WHERE id = $1`
	var mVideo Video
	err := c.client.QueryRow(ctx, query, videoID).Scan(&mVideo.ID, &mVideo.Url, &mVideo.Description, &mVideo.VectorIDs)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		} else {
			return nil, fmt.Errorf("failed to get video: %v", err)
		}
	}

	video := mVideo.ToEntity()
	return video, nil
}

func (c *Client) Update(ctx context.Context, video *entities.Video) error {
	query := `UPDATE videos SET (url, description, vectors) = ($2, $3, $4) WHERE id = $1`

	_, err := c.client.Exec(ctx, query, video.ID, video.URL, video.Description, video.RelatedVectorIDs)

	if err != nil {
		return fmt.Errorf("failed to update video: %v", err)
	}

	return nil
}
