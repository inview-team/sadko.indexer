package main

import (
	"context"
	"fmt"
	"os"

	"github.com/inview-team/sadko_indexer/config"
	"github.com/inview-team/sadko_indexer/internal/application/video"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/api"
)

const (
	videoServiceEnv = "SERVICE_CONFIG_PATH"
)

func main() {
	ctx := context.TODO()
	videoServiceConfigPath := os.Getenv(videoServiceEnv)
	if videoServiceConfigPath == "" {

		os.Exit(1)
	}

	cfg, err := config.LoadFile(videoServiceConfigPath)
	if err != nil {
		os.Exit(1)
	}

	app, err := video.NewApp(ctx, cfg.PostgresConfig, cfg.RabbitConfig)
	if err != nil {
		fmt.Print(err)
	}

	srv := api.NewServer(app)
	srv.Start(ctx)
}
