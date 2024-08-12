package harmony

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Driver(host string, path string, method string, body io.ReadCloser) (*http.Response, error) {
	cl := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
	}

	fullURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	req, err := http.NewRequest(method, fullURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to %v: %w", fullURL, err)
	}

	defer res.Body.Close()

	return res, nil
}
