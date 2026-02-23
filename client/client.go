package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	apiToken   string
	httpClient *http.Client
}

func NewClient(hostname string, apiToken string) *Client {
	client := &Client{
		hostname:   hostname,
		apiToken:   apiToken,
		httpClient: &http.Client{},
	}

	return client
}

// HttpRequest performs an HTTP request.
// Deprecated: Use HttpRequestWithContext instead.
func (client *Client) HttpRequest(path, method string, body ...bytes.Buffer) (closer io.ReadCloser, err error) {
	return client.HttpRequestWithContext(context.Background(), path, method, body...)
}

// HttpRequestWithContext performs an HTTP request with context support.
func (client *Client) HttpRequestWithContext(ctx context.Context, path, method string, body ...bytes.Buffer) (closer io.ReadCloser, err error) {
	url := client.requestPath(path)
	var bodyBuffer bytes.Buffer

	if len(body) > 0 {
		bodyBuffer = body[0]
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &bodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	if bodyBuffer.Len() > 0 {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthenticated")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid token")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}

	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return c.hostname + "/api/v1/" + path
}

func DecodeResponse[T any](body io.ReadCloser, target *T) (*T, error) {
	err := json.NewDecoder(body).Decode(target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

func EncodeRequest[T any](target *T) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(target)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
