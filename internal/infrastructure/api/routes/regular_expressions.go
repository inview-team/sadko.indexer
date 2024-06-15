package routes

import "fmt"

const (
	videoID = "video_id"
)

const regexp = "[a-f\\d]{8}-[a-f\\d]{4}-4[a-f\\d]{3}-[89ab][a-f\\d]{3}-[a-f\\d]{12}"

var patternVideoID = fmt.Sprintf("{%s:%s}", videoID, regexp)
