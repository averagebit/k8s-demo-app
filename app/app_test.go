package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/averagebit/k8s-demo-app/app"
)

func TestGreet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	app.Server(res, req)

	actual := res.Body.String()
	expected := "Hello from Go!"

	if expected != actual {
		t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
	}
}
