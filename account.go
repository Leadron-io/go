package leadron

import (
	"context"
	"net/url"
)

// AccountResource handles account and API key operations.
type AccountResource struct {
	client *Client
}

func (r *AccountResource) Get(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/account", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) Update(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PATCH", "/v1/account", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) GetBranding(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/account/branding", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) UpdateBranding(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/account/branding", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// API Keys

func (r *AccountResource) ApiKeysList(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/api-keys", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) ApiKeysCreate(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/api-keys", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) ApiKeysRevoke(ctx context.Context, keyID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/api-keys/"+url.PathEscape(keyID), nil, nil, opts)
	return err
}

func (r *AccountResource) GetUsage(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/account/usage", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) GetPlan(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/account/plan", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AccountResource) GetLimits(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/account/limits", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
