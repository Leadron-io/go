package leadron

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultBaseURL   = "https://api.leadron.io"
	defaultRetries   = 3
	retryDelaySec    = 1
	rateLimitHeader  = "X-RateLimit-Remaining"
	retryAfterHeader = "Retry-After"
)

// Client is the shared HTTP client for the SDK.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	maxRetries int
	rateLimit  *int
}

// RequestOptions holds optional headers for a request.
type RequestOptions struct {
	RequestID     string
	IdempotencyKey string
}

// Response holds data and optional pagination from an API response.
type Response struct {
	Data             interface{}
	Pagination       interface{}
	RateLimitRemaining *int
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body interface{}, opts *RequestOptions) (*Response, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = u.Path + path
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(raw)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if opts != nil {
		if opts.RequestID != "" {
			req.Header.Set("X-Request-Id", opts.RequestID)
		}
		if opts.IdempotencyKey != "" {
			req.Header.Set("Idempotency-Key", opts.IdempotencyKey)
		}
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(time.Duration(retryDelaySec*(attempt+1)) * time.Second)
				continue
			}
			return nil, err
		}

		var rateLimitRemaining *int
		if v := resp.Header.Get(rateLimitHeader); v != "" {
			if n, e := strconv.Atoi(v); e == nil && n >= 0 {
				rateLimitRemaining = &n
				if c.rateLimit != nil {
					*c.rateLimit = n
				}
			}
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var out struct {
				Data       json.RawMessage `json:"data"`
				Pagination json.RawMessage `json:"pagination"`
			}
			_ = json.Unmarshal(bodyBytes, &out)
			var data interface{}
			if len(out.Data) > 0 {
				_ = json.Unmarshal(out.Data, &data)
			} else {
				data = map[string]interface{}{}
			}
			return &Response{
				Data:             data,
				Pagination:       out.Pagination,
				RateLimitRemaining: rateLimitRemaining,
			}, nil
		}

		message := resp.Status
		var responseBody interface{}
		if len(bodyBytes) > 0 {
			_ = json.Unmarshal(bodyBytes, &responseBody)
			if m, ok := responseBody.(map[string]interface{}); ok && m["message"] != nil {
				if s, ok := m["message"].(string); ok {
					message = s
				}
			}
		}

		apiErr := &Error{Message: message, StatusCode: resp.StatusCode, Body: responseBody}

		switch resp.StatusCode {
		case 401:
			return nil, &AuthError{apiErr}
		case 422:
			return nil, &ValidationError{apiErr}
		case 429:
			retryAfter := retryDelaySec
			if v := resp.Header.Get(retryAfterHeader); v != "" {
				if n, e := strconv.Atoi(v); e == nil {
					retryAfter = n
				}
			}
			if attempt < c.maxRetries {
				time.Sleep(time.Duration(retryAfter) * time.Second)
				continue
			}
			return nil, &RateLimitError{apiErr, RetryAfter: retryAfter}
		}

		if resp.StatusCode >= 500 && resp.StatusCode < 600 && attempt < c.maxRetries {
			lastErr = apiErr
			time.Sleep(time.Duration(retryDelaySec*(attempt+1)) * time.Second)
			continue
		}
		return nil, apiErr
	}
	return nil, lastErr
}

// GetRateLimitStatus returns the last X-RateLimit-Remaining value, if any.
func (c *Client) GetRateLimitStatus() *int {
	return c.rateLimit
}
