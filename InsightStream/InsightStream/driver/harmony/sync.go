package harmony

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Driver(host string, path string, method string, body io.ReadCloser) (http.Response, error) {
	cl := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
	}

	targetURL, err := url.JoinPath(host, path)
	if err != nil {
		return http.Response{
			StatusCode: 500,
			Status:     "Internal Server Error",
		}, err
	}

	ul, err := url.Parse(targetURL)
	if err != nil {
		return http.Response{
			StatusCode: 500,
			Status:     "Internal Server Error",
		}, err
	}

	req := http.Request{
		Method: method,
		URL: &url.URL{
			Scheme: "http",
			Host:   host,
			Path:   path,
		},
		Body: body,
	}

	res, err := cl.Do(&req)
	if err != nil {
		return http.Response{
			StatusCode: 500,
			Status:     fmt.Sprintf("Internal Server Error in %v: %v", ul, err.Error()),
		}, err
	}

	return *res, nil
}
