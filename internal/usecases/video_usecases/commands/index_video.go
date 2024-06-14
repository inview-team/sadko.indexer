package commands

import (
	"context"
	"fmt"

	"github.com/inview-team/sadko_indexer/internal/entities"
)

type IndexVideoCommand struct {
	repo entities.VideoRepository
}

func NewIndexVideoCommand(repo entities.VideoRepository) IndexVideoCommand {
	return IndexVideoCommand{repo}
}

func (c *IndexVideoCommand) Execute(ctx context.Context, url string, description string) (string, error) {
	video := entities.NewVideo(c.repo.NextID(ctx), url, description, []entities.VectorID{})
	err := c.repo.Create(ctx, video)
	if err != nil {
		fmt.Printf("failed to index video: %v", err)
		return "", err
	}
	return string(video.ID), nil
}
