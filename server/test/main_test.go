package test

import (
	"net/http"
	"net/http/httptest"
	"serveur/server/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnauthentifiedUserCannotAccessToProtectedRoute(t *testing.T) {
	router := router.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/authtest", nil)
	router.ServeHTTP(w, req)

	// 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthentifiedUserCanAccessToProtectedRoute(t *testing.T) {
	// user mock objects

	router := router.SetUpRouter()
}
