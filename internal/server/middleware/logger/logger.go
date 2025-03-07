package logger

import (
	"log/slog"
	"net/http"
	"time"
)

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = log.With(slog.String("component", "middleware/logger"))
		log.Info("logger middleware enabled")
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.RequestURI),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", r.Header.Get("X-Request-Id")),
			)

			t1 := time.Now()
			defer func() {
				entry.Info("request",
					slog.Int("status", r.Response.StatusCode),
					slog.Int("bytes", int(r.ContentLength)),
					slog.String("duration", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
