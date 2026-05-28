package leadron

import (
	"context"
	"net/url"
)

// PartnersResource handles partner operations.
type PartnersResource struct {
	client *Client
}

func (r *PartnersResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/partners", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) Get(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) Update(ctx context.Context, partnerID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/partners/"+url.PathEscape(partnerID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) List(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) Deactivate(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/partners/"+url.PathEscape(partnerID)+"/deactivate", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetReferralTree(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/hierarchy", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetUpline(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/hierarchy", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetReferralLink(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/referral-link", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetStats(ctx context.Context, partnerID string, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/metrics", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetLeaderboard(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/leaderboard", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetTopPerformers(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	return r.GetLeaderboard(ctx, params, opts)
}

func (r *PartnersResource) Invite(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/partners/invite", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) ResendInvite(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/partners/"+url.PathEscape(partnerID)+"/resend-invite", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetOnboardingStatus(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/onboarding", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) SendAgreement(ctx context.Context, partnerID, templateID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/partners/"+url.PathEscape(partnerID)+"/agreements/send", nil, map[string]string{"templateId": templateID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetSignedDocuments(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/documents/signed", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PartnersResource) GetAgreementStatus(ctx context.Context, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/partners/"+url.PathEscape(partnerID)+"/agreement-status", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
