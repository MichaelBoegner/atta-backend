package main

import (
	"net/http/httptest"
	"testing"
)

func TestGetMessage(t *testing.T) {
	t.Run("should handle returning messages from a third party source", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()

		MessageServer(response, request)

		got := response.Body.String()
		want := "He did a great job today. Attaboy!"

		if got != want {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})
	t.Run("get message from test endpoint")
}
