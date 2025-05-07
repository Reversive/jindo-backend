package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	t.Run("allowed method", func(t *testing.T) {
		r := NewRouter()
		r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/get", nil)
		rr := httptest.NewRecorder()
		r.Handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v but want %v",
				status, http.StatusOK)
		}
	})

	t.Run("not allowed method", func(t *testing.T) {
		r := NewRouter()

		r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
		})

		req, _ := http.NewRequest(http.MethodPost, "/get", nil)
		rr := httptest.NewRecorder()

		r.Handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("request returned wrong status code, got %v but want %v",
				status, http.StatusMethodNotAllowed)
		}
	})
}
