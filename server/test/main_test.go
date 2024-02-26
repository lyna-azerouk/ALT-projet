package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	router2 "serveur/server/router"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := router2.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
