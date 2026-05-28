package leadron

import (
	"context"
	"net/url"
)

// AnalyticsResource handles analytics operations.
type AnalyticsResource struct {
	client *Client
}

func (r *AnalyticsResource) GetOverview(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/analytics/overview", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AnalyticsResource) GetLeadMetrics(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/analytics/leads", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AnalyticsResource) GetCommissionMetrics(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/analytics/commissions", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AnalyticsResource) GetPartnerMetrics(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/analytics/partners", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AnalyticsResource) GetConversionRate(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/analytics/conversion-rate", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *AnalyticsResource) GetSmsMetrics(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/analytics/sms", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
