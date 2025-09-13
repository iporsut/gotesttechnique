package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func SendRequestWithContext(ctx context.Context, url string, method string, body io.Reader) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
