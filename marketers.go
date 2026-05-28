package leadron

import (
	"context"
	"net/url"
)

// MarketersResource handles marketer operations.
type MarketersResource struct {
	client *Client
}

func (r *MarketersResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/marketers", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) Get(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) Update(ctx context.Context, partnerID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/marketers/"+url.PathEscape(partnerID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) List(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) Deactivate(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/marketers/"+url.PathEscape(partnerID)+"/deactivate", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetReferralTree(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/hierarchy", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetUpline(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/hierarchy", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetReferralLink(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/referral-link", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetStats(ctx context.Context, partnerID string, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/metrics", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetLeaderboard(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/leaderboard", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetTopPerformers(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	return r.GetLeaderboard(ctx, params, opts)
}

func (r *MarketersResource) Invite(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/marketers/invite", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) ResendInvite(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/marketers/"+url.PathEscape(partnerID)+"/resend-invite", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetOnboardingStatus(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/onboarding", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) SendAgreement(ctx context.Context, partnerID, templateID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/marketers/"+url.PathEscape(partnerID)+"/agreements/send", nil, map[string]string{"templateId": templateID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetSignedDocuments(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/documents/signed", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *MarketersResource) GetAgreementStatus(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/marketers/"+url.PathEscape(partnerID)+"/agreement-status", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
