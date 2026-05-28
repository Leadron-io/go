package leadron

import (
	"context"
	"net/url"
)

// PhoneNumbersResource handles phone number operations.
type PhoneNumbersResource struct {
	client *Client
}

func (r *PhoneNumbersResource) Search(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/phone-numbers", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PhoneNumbersResource) List(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/phone-numbers", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PhoneNumbersResource) Get(ctx context.Context, numberID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/phone-numbers/"+url.PathEscape(numberID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PhoneNumbersResource) Release(ctx context.Context, numberID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/phone-numbers/"+url.PathEscape(numberID), nil, nil, opts)
	return err
}

func (r *PhoneNumbersResource) AssignToTeam(ctx context.Context, numberID, teamID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/phone-numbers/"+url.PathEscape(numberID)+"/assign-team", nil, map[string]string{"teamId": teamID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PhoneNumbersResource) UnassignFromTeam(ctx context.Context, numberID, teamID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/phone-numbers/"+url.PathEscape(numberID)+"/unassign-team", nil, map[string]string{"teamId": teamID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PhoneNumbersResource) GetUsage(ctx context.Context, numberID string, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/phone-numbers/"+url.PathEscape(numberID)+"/usage", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *PhoneNumbersResource) Get10DLCStatus(ctx context.Context, numberID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/phone-numbers/"+url.PathEscape(numberID)+"/10dlc-status", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
