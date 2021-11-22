package metadata

import (
	"io"
	"net/http"
)

// NewRequest realize the request to link
func NewRequest(url string) (res io.ReadCloser, err error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	res = resp.Body
	return
}
