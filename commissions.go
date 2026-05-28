package leadron

import (
	"context"
	"net/url"
)

// CommissionsResource handles commission operations.
type CommissionsResource struct {
	client *Client
}

func (r *CommissionsResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/commissions/records", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) Get(ctx context.Context, commissionID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/records/"+url.PathEscape(commissionID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) List(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/records", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) Approve(ctx context.Context, commissionID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/commissions/records/"+url.PathEscape(commissionID)+"/approve", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) Reject(ctx context.Context, commissionID, reason string, opts *RequestOptions) (interface{}, error) {
	body := interface{}(nil)
	if reason != "" {
		body = map[string]string{"reason": reason}
	}
	resp, err := r.client.do(ctx, "POST", "/v1/commissions/records/"+url.PathEscape(commissionID)+"/reject", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) MarkPaid(ctx context.Context, commissionID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/commissions/records/"+url.PathEscape(commissionID)+"/mark-paid", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) GetRules(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/rules", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) CreateRule(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/commissions/rules", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) UpdateRule(ctx context.Context, ruleID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/commissions/rules/"+url.PathEscape(ruleID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) DeleteRule(ctx context.Context, ruleID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/commissions/rules/"+url.PathEscape(ruleID), nil, nil, opts)
	return err
}

func (r *CommissionsResource) GetPayoutSummary(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/partners/"+url.PathEscape(partnerID)+"/payout-summary", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) GetWalletBalance(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/partners/"+url.PathEscape(partnerID)+"/wallet", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) RequestPayout(ctx context.Context, partnerID string, amount float64, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/commissions/partners/"+url.PathEscape(partnerID)+"/payouts", nil, map[string]float64{"amount": amount}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) GetPayoutHistory(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/partners/"+url.PathEscape(partnerID)+"/payouts", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) GetSummary(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/summary", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) GetByPartner(ctx context.Context, partnerID string, params url.Values, opts *RequestOptions) (interface{}, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("partnerId", partnerID)
	return r.List(ctx, params, opts)
}

func (r *CommissionsResource) GetTotalOwed(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/total-owed", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *CommissionsResource) GetTotalPaid(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/commissions/total-paid", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
