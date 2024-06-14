package postgres

import "github.com/inview-team/sadko_indexer/internal/entities"

type Video struct {
	ID          string
	Url         string
	Description string
	VectorIDs   []string
}

func (v Video) ToEntity() *entities.Video {
	var vectors []entities.VectorID
	for _, app := range v.VectorIDs {
		vectors = append(vectors, entities.VectorID(app))
	}

	return entities.NewVideo(entities.VideoID(v.ID), v.Url, v.Description, vectors)
}
