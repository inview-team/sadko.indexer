package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		ctx := r.Context()
		req := r.WithContext(context.WithValue(ctx, "request_id", uuid.New().String()))
		*r = *req

		next.ServeHTTP(recorder, req)

	})
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
