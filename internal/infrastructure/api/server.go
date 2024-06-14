package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inview-team/sadko_indexer/internal/application/video"
	"github.com/inview-team/sadko_indexer/internal/infrastructure/api/routes"
)

type Server struct {
	srv http.Server
}

func NewServer(app *video.App) *Server {
	return &Server{
		srv: http.Server{
			Handler:      routes.Make(app),
			Addr:         ":30001",
			IdleTimeout:  time.Minute,
			ReadTimeout:  time.Minute,
			WriteTimeout: time.Minute,
		},
	}
}

func (s *Server) Start(ctx context.Context) {
	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		// Listen on application shutdown signals.

		// Shutdown HTTP server.
		if err := s.srv.Shutdown(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		}
	}()

	fmt.Println("Listening on ", s.srv.Addr)
	// Start HTTP server.
	if err := s.srv.ListenAndServe(); err != nil {
		fmt.Printf("Failed to listen and serve: %s", err)
	}
}
