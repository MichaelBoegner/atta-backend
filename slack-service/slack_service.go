package slack_service

import (
	"fmt"
	"io"
	"net/http"
)

func GETSlackMessages(name string) (message string) {
	requestURL := fmt.Sprint("https://webhook.site/56c3dcfb-1056-42e2-abf3-ae668cf14be4")
	res, _ := http.Get(requestURL)

	resBytes, _ := io.ReadAll(res.Body)

	return string(resBytes)
}
