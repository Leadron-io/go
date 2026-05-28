package leadron

import (
	"context"
	"net/url"
)

// TeamsResource handles team operations.
type TeamsResource struct {
	client *Client
}

func (r *TeamsResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/teams", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) Get(ctx context.Context, teamID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/teams/"+url.PathEscape(teamID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) List(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/teams", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) Update(ctx context.Context, teamID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/teams/"+url.PathEscape(teamID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) Delete(ctx context.Context, teamID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/teams/"+url.PathEscape(teamID), nil, nil, opts)
	return err
}

func (r *TeamsResource) AddMember(ctx context.Context, teamID, userID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/teams/"+url.PathEscape(teamID)+"/members", nil, map[string]string{"userId": userID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) RemoveMember(ctx context.Context, teamID, userID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/teams/"+url.PathEscape(teamID)+"/members/"+url.PathEscape(userID), nil, nil, opts)
	return err
}

func (r *TeamsResource) GetMembers(ctx context.Context, teamID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/teams/"+url.PathEscape(teamID)+"/members", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) AssignLead(ctx context.Context, teamID, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/teams/"+url.PathEscape(teamID)+"/assign/lead", nil, map[string]string{"leadId": leadID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) AssignPhoneNumber(ctx context.Context, teamID, numberID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/teams/"+url.PathEscape(teamID)+"/assign/phone-number", nil, map[string]string{"numberId": numberID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) GetStats(ctx context.Context, teamID string, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/teams/"+url.PathEscape(teamID)+"/stats", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *TeamsResource) GetLeaderboard(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/teams/leaderboard", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
