package leadron

import (
	"context"
	"net/url"
)

// ReportsResource handles report operations.
type ReportsResource struct {
	client *Client
}

func (r *ReportsResource) Leads(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/reports/leads", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *ReportsResource) Commissions(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/reports/commissions", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *ReportsResource) Partners(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/reports/partners", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *ReportsResource) Export(ctx context.Context, reportConfig interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/reports/export", nil, reportConfig, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
