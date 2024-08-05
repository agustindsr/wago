package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"wasm/client/wa/dom"
)

func Fetch[T any](urlStr string, queryParams map[string]string, headers map[string]string) dom.Func {
	return func(e dom.Event) {
		go func() {
			FetchData[T](urlStr, queryParams, headers)
		}()
	}
}

func FetchData[T any](urlStr string, queryParams map[string]string, headers map[string]string) (T, error) {
	var response T
	u, err := url.Parse(urlStr)
	if err != nil {
		return response, fmt.Errorf("error parsing URL: %v", err)
	}

	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return response, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Access-Control-Request-Method", "GET")
	req.Header.Set("Origin", "https://agustindsr.github.io")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("error parsing JSON: %v", err)
	}

	return response, nil
}
