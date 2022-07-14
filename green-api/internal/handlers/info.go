package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Info is a http.Handler
type Info struct {
	log *logrus.Entry
}

// NewInfo creates a info handler
func NewInfo(log *logrus.Entry) *Info {
	return &Info{
		log: log,
	}
}

// ServeHTTP is the main entry point for the handler
func (i *Info) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	i.log.WithField("endpoint", r.URL.Path).WithField("method", r.Method).Info("processing request.")

	rw.Write([]byte(fmt.Sprintf("Server time: %v", time.Now())))
}
