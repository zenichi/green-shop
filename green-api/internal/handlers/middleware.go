package handlers

import (
	"net/http"
	"time"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

type ApiMiddleware struct {
	log *logrus.Entry
}

func NewApiMiddleware(log *logrus.Entry) *ApiMiddleware {
	return &ApiMiddleware{log}
}

func (api *ApiMiddleware) WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log := api.log.WithFields(logrus.Fields{
			"uri":    r.RequestURI,
			"method": r.Method,
			"id":     xid.New(),
		})
		log.Info("request started")

		defer func(log *logrus.Entry) {
			if err := recover(); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				log.WithField("error", err).Error("internal server error.")
			}
		}(log)

		start := time.Now()

		wrapped := newWrappedResponseWriter(rw)
		next.ServeHTTP(wrapped, r)

		log.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"status":   wrapped.status,
		}).Info("request completed")
	})
}

// wrappedResponseWriter is a wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type wrappedResponseWriter struct {
	http.ResponseWriter
	status int
}

func newWrappedResponseWriter(wr http.ResponseWriter) *wrappedResponseWriter {
	return &wrappedResponseWriter{ResponseWriter: wr}
}

func (w *wrappedResponseWriter) Status() int {
	return w.status
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
