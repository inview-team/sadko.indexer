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
	playbookServiceENV = "SERVICE_CONFIG_PATH"
)

func main() {
	ctx := context.TODO()
	playbookServiceConfigPath := os.Getenv(playbookServiceENV)
	if playbookServiceConfigPath == "" {

		os.Exit(1)
	}

	cfg, err := config.LoadFile(playbookServiceConfigPath)
	if err != nil {
		os.Exit(1)
	}

	app, err := video.NewApp(ctx, cfg.PostgresConfig)
	if err != nil {
		fmt.Print(err)
	}

	srv := api.NewServer(app)
	srv.Start(ctx)
}
