package leadron

import (
	"context"
	"encoding/json"
	"net/url"
)

// WebhooksResource handles webhook operations.
type WebhooksResource struct {
	client *Client
}

func (r *WebhooksResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/integrations/webhooks", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *WebhooksResource) List(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/integrations/webhooks", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *WebhooksResource) Get(ctx context.Context, webhookID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/integrations/webhooks/"+url.PathEscape(webhookID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *WebhooksResource) Update(ctx context.Context, webhookID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/integrations/webhooks/"+url.PathEscape(webhookID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *WebhooksResource) Delete(ctx context.Context, webhookID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/integrations/webhooks/"+url.PathEscape(webhookID), nil, nil, opts)
	return err
}

func (r *WebhooksResource) Test(ctx context.Context, webhookID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/integrations/webhooks/"+url.PathEscape(webhookID)+"/test", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *WebhooksResource) GetLogs(ctx context.Context, webhookID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/integrations/webhooks/"+url.PathEscape(webhookID)+"/logs", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *WebhooksResource) Retry(ctx context.Context, webhookID, logID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/integrations/webhooks/"+url.PathEscape(webhookID)+"/retry", nil, map[string]string{"logId": logID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// ConstructEvent verifies the webhook signature and parses the JSON payload. Returns the decoded event or an error.
func ConstructEvent(rawBody, signature, secret string) (interface{}, error) {
	if !VerifyWebhookSignature(rawBody, signature, secret) {
		return nil, &Error{Message: "webhook signature verification failed"}
	}
	var out interface{}
	if err := json.Unmarshal([]byte(rawBody), &out); err != nil {
		return nil, err
	}
	return out, nil
}
