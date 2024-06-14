package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inview-team/sadko_indexer/internal/application/video"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/api/controllers"
	"github.com/inview-team/sadko_indexer/internal/usecases/video_usecases"
)

func indexVideo(usecases video_usecases.VideoUsecases) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error index video"
		ctx := r.Context()

		cVideo := new(controllers.IndexVideoPayload)
		if err := json.NewDecoder(r.Body).Decode(cVideo); err != nil {
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		_, err := usecases.IndexVideo.Execute(ctx, cVideo.Url, cVideo.Description)

		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}
	})
}

func makeVideoRoutes(r *mux.Router, app *video.App) {
	path := "/index"
	serviceRouter := r.PathPrefix(path).Subrouter()
	serviceRouter.Handle("", indexVideo(app.Video)).Methods("POST")
}
