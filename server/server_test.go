package server

import "testing"

func TestServer(t *testing.T) {
	t.Run("should GET messages from Slack", func(t *testing.T) {
		got := Server("Mike")
		want := "He did a great job today. Attaboy!"

		if got != want {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})
}
