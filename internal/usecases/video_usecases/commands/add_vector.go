package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/inview-team/sadko_indexer/internal/entities"
)

type AddVectorsCommand struct {
	repo entities.VideoRepository
}

func NewAddVectorCommand(repo entities.VideoRepository) AddVectorsCommand {
	return AddVectorsCommand{repo}
}

func (c *AddVectorsCommand) Execute(ctx context.Context, videoID string, vectors []entities.VectorID) error {
	video, err := c.repo.GetByID(ctx, videoID)
	if err != nil {
		return fmt.Errorf("failed to add vectors: %v", err)
	}

	if video == nil {
		return ErrVideoNotFound
	}

	video.RelatedVectorIDs = vectors

	err = c.repo.Update(ctx, video)
	if err != nil {
		return fmt.Errorf("failed to add vectors: %v", err)
	}
	return nil
}

var (
	ErrVideoNotFound = errors.New("video does not exist")
)
