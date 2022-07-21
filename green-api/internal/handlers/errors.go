package handlers

import (
	"log"
	"net/http"

	"github.com/zenichi/green-shop/green-api/internal/utils"
)

// errorResponse is a generic API error structure
type errorResponse struct {
	Message string `json:"message"`
}

// genericErrorResponse sets the response header and returns generic response in JSON format
func genericErrorResponse(rw http.ResponseWriter, statusCode int, errMsg string) {
	rw.WriteHeader(statusCode)
	err := utils.ToJSON(&errorResponse{Message: errMsg}, rw)
	if err != nil {
		log.Fatal("serialization error")
	}
}

type validationErrors struct {
	*errorResponse
	ValidationMessages []string `json:"validationMessages"`
}

func validationErrorsResponse(rw http.ResponseWriter, errors []string) {
	rw.WriteHeader(http.StatusUnprocessableEntity)
	err := utils.ToJSON(&validationErrors{&errorResponse{"Invalid data"}, errors}, rw)
	if err != nil {
		log.Fatal("serialization error")
	}
}
