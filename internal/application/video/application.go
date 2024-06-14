package video

import (
	"context"
	"fmt"

	"github.com/inview-team/sadko_indexer/internal/infrastructure/video_repository/postgres"
	"github.com/inview-team/sadko_indexer/internal/usecases/video_usecases"
)

type App struct {
	Video video_usecases.VideoUsecases
}

func NewApp(ctx context.Context, config postgres.Config) (*App, error) {
	pgClient, err := postgres.New(ctx, config)
	if err != nil {
		fmt.Printf("failed to create application: %v", err)
		return nil, err
	}

	app := App{
		Video: video_usecases.NewVideoUsecases(pgClient),
	}

	return &app, nil
}
