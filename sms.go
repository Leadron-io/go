package leadron

import (
	"context"
	"net/url"
)

// SMSResource handles SMS/communications operations.
type SMSResource struct {
	client *Client
}

func (r *SMSResource) Send(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/communications/send/sms", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SMSResource) GetInbox(ctx context.Context, phoneNumberID string, opts *RequestOptions) (interface{}, error) {
	q := url.Values{}
	if phoneNumberID != "" {
		q.Set("phoneNumberId", phoneNumberID)
	}
	resp, err := r.client.do(ctx, "GET", "/v1/communications/inbox", q, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SMSResource) GetOutbox(ctx context.Context, phoneNumberID string, opts *RequestOptions) (interface{}, error) {
	q := url.Values{}
	if phoneNumberID != "" {
		q.Set("phoneNumberId", phoneNumberID)
	}
	resp, err := r.client.do(ctx, "GET", "/v1/communications/outbox", q, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SMSResource) GetConversation(ctx context.Context, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/communications/conversations/"+url.PathEscape(leadID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SMSResource) GetUsage(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/communications/usage", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
