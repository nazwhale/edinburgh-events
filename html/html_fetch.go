package html

import (
	"io"
	"net/http"
)

// FetchHTML makes an HTTP GET request to the specified URL and returns the HTML content.
func FetchHTML(url string) (string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
