package video

import (
	"context"
	"fmt"

	"github.com/inview-team/sadko_indexer/internal/infrastructure/tag_service/rabbitmq"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/video_repository/postgres"
	"github.com/inview-team/sadko_indexer/internal/usecases/video_usecases"
)

type App struct {
	Video video_usecases.VideoUsecases
}

func NewApp(ctx context.Context, pConfig postgres.Config, rConfig rabbitmq.Config) (*App, error) {
	pgClient, err := postgres.New(ctx, pConfig)
	if err != nil {
		fmt.Printf("failed to create application: %v", err)
		return nil, err
	}

	rmq, err := rabbitmq.New(rConfig)
	if err != nil {
		fmt.Printf("failed to create application: %v", err)
		return nil, err
	}

	app := App{
		Video: video_usecases.NewVideoUsecases(pgClient, rmq),
	}

	return &app, nil
}
