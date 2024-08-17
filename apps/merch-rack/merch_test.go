package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

// func TestMain(m *testing.M) {
// 	m.Run()
// }

func TestGetMerchRoute(t *testing.T) {
	r := getRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/merch/", nil)
	r.ServeHTTP(w, req)

	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
