package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/D3vR4pt0rs/logger"
	"github.com/gorilla/mux"
	"github.com/inview-team/sadko_indexer/internal/application/video"
	"github.com/inview-team/sadko_indexer/internal/entities"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/api/controllers"
	"github.com/inview-team/sadko_indexer/internal/usecases/video_usecases"
	"github.com/inview-team/sadko_indexer/internal/usecases/video_usecases/commands"
)

func indexVideo(usecases video_usecases.VideoUsecases) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error index video"
		ctx := r.Context()

		cVideo := new(controllers.IndexVideoPayload)
		if err := json.NewDecoder(r.Body).Decode(cVideo); err != nil {
			logger.Error.Print("failed to decode payload")
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		_, err := usecases.IndexVideo.Execute(ctx, cVideo.Url, cVideo.Description)

		if err != nil {
			logger.Error.Printf("failed to index video: %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}
	})
}

func addVectorsId(usecases video_usecases.VideoUsecases) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error add vectors id to video"

		ctx := r.Context()

		mVectors := new(controllers.VectorsPayload)
		if err := json.NewDecoder(r.Body).Decode(mVectors); err != nil {
			logger.Error.Print("failed to decode payload: %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
		}

		var vectors []entities.VectorID
		for _, v := range mVectors.Vectors {
			vectors = append(vectors, entities.VectorID(v))
		}

		err := usecases.AddVectors.Execute(ctx, mux.Vars(r)[videoID], vectors)
		if err != nil {
			logger.Error.Print("failed to decode payload: %v", err)
			if errors.Is(err, commands.ErrVideoNotFound) {
				http.Error(w, errorMessage, http.StatusNotFound)
				return
			}

			http.Error(w, errorMessage, http.StatusInternalServerError)
		}
	})
}

func makeVideoRoutes(r *mux.Router, app *video.App) {
	path := "/index"
	serviceRouter := r.PathPrefix(path).Subrouter()
	serviceRouter.Handle("", indexVideo(app.Video)).Methods("POST")
	serviceRouter.Handle(fmt.Sprintf("/%s", patternVideoID), addVectorsId(app.Video)).Methods("POST")
}
