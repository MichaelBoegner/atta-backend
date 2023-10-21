package slack_service

import "testing"

func TestGETSlackMessages(t *testing.T) {
	t.Run("should get messages from test link mocking slack")
	got := GETSlackMessages("Mike")
	want := "Michael is doing a great job. Attaboy!"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
