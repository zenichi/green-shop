package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInfoReturnsValidStatusCode(t *testing.T) {
	ih := NewInfo(logrus.WithField("context", "tests"))

	r := httptest.NewRequest(http.MethodGet, "/info", nil)
	response := httptest.NewRecorder()
	ih.ServeHTTP(response, r)

	assert.Equal(t, response.Code, http.StatusOK, "status should be 200")
}
