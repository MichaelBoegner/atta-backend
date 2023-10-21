package slack_service

import (
	"testing"
)

func TestGETSlackMessages(t *testing.T) {
	t.Run("should get messages from test link mocking slack", func(t *testing.T) {
		got := GETSlackMessages("Mike")
		want := "\"Michael is doing a great job. Attaboy!\""

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

}
