package entities

import "context"

type Video struct {
	ID               VideoID
	URL              string
	Description      string
	RelatedVectorIDs []VectorID
}

type VideoID string

type VectorID string

func NewVideo(id VideoID, url string, desc string, vectorIDs []VectorID) *Video {
	return &Video{
		ID:               id,
		URL:              url,
		Description:      desc,
		RelatedVectorIDs: vectorIDs,
	}
}

type VideoRepository interface {
	Create(ctx context.Context, video *Video) error
	GetByID(ctx context.Context, videoID string) (*Video, error)
	Update(ctx context.Context, video *Video) error
	NextID(ctx context.Context) VideoID
}
