package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-shop/green-api/internal/data"
	"github.com/zenichi/green-shop/green-api/internal/utils"
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

func (ph *Product) ValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p := &data.Product{}

		err := utils.FromJSON(p, r.Body)
		if err != nil {
			ph.log.WithError(err).Error("Unable to deserialize from JSON")
			genericErrorResponse(rw, http.StatusBadRequest, "Product has invalid structure.")
			return
		}

		errors := ph.v.Validate(p)
		if len(errors) > 0 {
			validationErrorsResponse(rw, errors)
			return
		}

		// store product in context, so that handler can retrieve it
		// without a need to unmarshal/validate request again
		ctx := context.WithValue(r.Context(), KeyDataProduct{}, p)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
