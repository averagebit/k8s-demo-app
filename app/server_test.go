package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/averagebit/circle-docker/api/app"
)

func TestGetRoot(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	app.Server(res, req)

	actual := res.Body.String()
	expected := "Hello World!"

	if actual != expected {
		t.Errorf("got %q, want %q", actual, expected)
	}
}
