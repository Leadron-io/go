package leadron

import (
	"context"
	"net/url"
	"strconv"
)

// LeadsResource handles lead operations.
type LeadsResource struct {
	client *Client
}

// Create creates a lead.
func (r *LeadsResource) Create(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Get returns a lead by ID.
func (r *LeadsResource) Get(ctx context.Context, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/leads/"+url.PathEscape(leadID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Update updates a lead.
func (r *LeadsResource) Update(ctx context.Context, leadID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/leads/"+url.PathEscape(leadID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Delete deletes a lead.
func (r *LeadsResource) Delete(ctx context.Context, leadID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/leads/"+url.PathEscape(leadID), nil, nil, opts)
	return err
}

// ListParams holds query params for listing leads.
type ListParams struct {
	Page        int
	Limit       int
	Status      string
	AssignedTo  string
	Source      string
	Sort        string
	Order       string
	From        string
	To          string
}

// List returns a paginated list of leads. Pagination and Data are in the response.
func (r *LeadsResource) List(ctx context.Context, params *ListParams, opts *RequestOptions) (data interface{}, pagination interface{}, err error) {
	q := url.Values{}
	if params != nil {
		if params.Page > 0 {
			q.Set("page", strconv.Itoa(params.Page))
		}
		if params.Limit > 0 {
			q.Set("limit", strconv.Itoa(params.Limit))
		}
		if params.Status != "" {
			q.Set("status", params.Status)
		}
		if params.AssignedTo != "" {
			q.Set("assignedTo", params.AssignedTo)
		}
		if params.Source != "" {
			q.Set("source", params.Source)
		}
		if params.Sort != "" {
			q.Set("sort", params.Sort)
		}
		if params.Order != "" {
			q.Set("order", params.Order)
		}
		if params.From != "" {
			q.Set("from", params.From)
		}
		if params.To != "" {
			q.Set("to", params.To)
		}
	}
	resp, err := r.client.do(ctx, "GET", "/v1/leads", q, nil, opts)
	if err != nil {
		return nil, nil, err
	}
	return resp.Data, resp.Pagination, nil
}

// Assign assigns a lead to a user (partner).
func (r *LeadsResource) Assign(ctx context.Context, leadID, userID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads/"+url.PathEscape(leadID)+"/assign", nil, map[string]string{"userId": userID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// UpdateStatus updates a lead's status.
func (r *LeadsResource) UpdateStatus(ctx context.Context, leadID, status string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/leads/"+url.PathEscape(leadID)+"/status", nil, map[string]string{"status": status}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// AddNote adds a note to a lead.
func (r *LeadsResource) AddNote(ctx context.Context, leadID, content string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads/"+url.PathEscape(leadID)+"/notes", nil, map[string]string{"content": content}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// GetNotes returns notes for a lead.
func (r *LeadsResource) GetNotes(ctx context.Context, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/leads/"+url.PathEscape(leadID)+"/notes", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// GetTimeline returns the timeline for a lead.
func (r *LeadsResource) GetTimeline(ctx context.Context, leadID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/leads/"+url.PathEscape(leadID)+"/timeline", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// MarkConverted marks a lead as converted. Opts can include dealValue, closedAt, notes.
func (r *LeadsResource) MarkConverted(ctx context.Context, leadID string, opts interface{}, reqOpts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads/"+url.PathEscape(leadID)+"/convert", nil, opts, reqOpts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// BulkCreate creates multiple leads.
func (r *LeadsResource) BulkCreate(ctx context.Context, leads interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads/bulk", nil, map[string]interface{}{"leads": leads}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// BulkAssign assigns multiple leads to a partner.
func (r *LeadsResource) BulkAssign(ctx context.Context, leadIDs []string, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads/bulk/assign", nil, map[string]interface{}{"leadIds": leadIDs, "partnerId": partnerID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// BulkUpdateStatus updates status for multiple leads.
func (r *LeadsResource) BulkUpdateStatus(ctx context.Context, leadIDs []string, status string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/leads/bulk/status", nil, map[string]interface{}{"leadIds": leadIDs, "status": status}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Search searches leads.
func (r *LeadsResource) Search(ctx context.Context, q string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/search/leads", url.Values{"q": {q}}, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Filter lists leads with filter params.
func (r *LeadsResource) Filter(ctx context.Context, params *ListParams, opts *RequestOptions) (interface{}, error) {
	data, _, err := r.List(ctx, params, opts)
	return data, err
}
