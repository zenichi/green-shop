package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-shop/green-api/internal/utils"
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
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	err := utils.ToJSON(infoResponse{Message: fmt.Sprintf("Server time: %v", time.Now())}, rw)
	if err != nil {
		i.log.Fatal("serialization error.")
	}
}

type infoResponse struct {
	Message string
}
