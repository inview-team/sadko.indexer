package rabbitmq

import "github.com/inview-team/sadko_indexer/internal/entities"

type TagMessage struct {
	ID          string `json:"id"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func NewTagMessage(video *entities.Video) *TagMessage {
	return &TagMessage{
		ID:          string(video.ID),
		Url:         video.URL,
		Description: video.Description,
	}
}
