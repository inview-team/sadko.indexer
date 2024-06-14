package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inview-team/sadko_indexer/internal/application/video"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/api/handlers"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/api/middleware"
)

func Make(app *video.App) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.AccessLogMiddleware)
	r.MethodNotAllowedHandler = handlers.NotAllowedHandler()
	r.NotFoundHandler = handlers.NotFoundHandler()

	makeVideoRoutes(r, app)
	return r
}
