package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	t.Parallel()

	t.Run("returns 200 OK", func(t *testing.T) {
		t.Parallel()

		request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		response := httptest.NewRecorder()

		NewRouter().ServeHTTP(response, request)

		if got, want := response.Code, http.StatusOK; got != want {
			t.Errorf("status = %d, want %d", got, want)
		}
	})
}
