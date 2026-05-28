package leadron

import (
	"context"
	"net/url"
)

// SequencesResource handles sequence operations.
type SequencesResource struct {
	client *Client
}

func (r *SequencesResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) Get(ctx context.Context, sequenceID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/sequences/"+url.PathEscape(sequenceID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) List(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/sequences", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) Update(ctx context.Context, sequenceID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/sequences/"+url.PathEscape(sequenceID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) Delete(ctx context.Context, sequenceID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/sequences/"+url.PathEscape(sequenceID), nil, nil, opts)
	return err
}

func (r *SequencesResource) Activate(ctx context.Context, sequenceID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/activate", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) Pause(ctx context.Context, sequenceID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/pause", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) EnrollLead(ctx context.Context, sequenceID, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/enroll", nil, map[string]string{"leadId": leadID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) EnrollBulk(ctx context.Context, sequenceID string, leadIDs []string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/enroll/bulk", nil, map[string]interface{}{"leadIds": leadIDs}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) UnenrollLead(ctx context.Context, sequenceID, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/unenroll/"+url.PathEscape(leadID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) GetEnrolledLeads(ctx context.Context, sequenceID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/sequences/"+url.PathEscape(sequenceID)+"/enrolled", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) AddStep(ctx context.Context, sequenceID string, step interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/steps", nil, step, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) UpdateStep(ctx context.Context, sequenceID, stepID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/sequences/"+url.PathEscape(sequenceID)+"/steps/"+url.PathEscape(stepID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *SequencesResource) DeleteStep(ctx context.Context, sequenceID, stepID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/sequences/"+url.PathEscape(sequenceID)+"/steps/"+url.PathEscape(stepID), nil, nil, opts)
	return err
}

func (r *SequencesResource) ReorderSteps(ctx context.Context, sequenceID string, stepOrder []string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/sequences/"+url.PathEscape(sequenceID)+"/steps/reorder", nil, map[string]interface{}{"stepOrder": stepOrder}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
