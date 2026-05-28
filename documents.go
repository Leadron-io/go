package leadron

import (
	"context"
	"net/url"
)

// DocumentsResource handles document and template operations.
type DocumentsResource struct {
	client *Client
}

// Templates

func (r *DocumentsResource) TemplatesCreate(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/documents/templates", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) TemplatesList(ctx context.Context, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents/templates", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) TemplatesGet(ctx context.Context, templateID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents/templates/"+url.PathEscape(templateID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) TemplatesUpdate(ctx context.Context, templateID string, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "PUT", "/v1/documents/templates/"+url.PathEscape(templateID), nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) TemplatesDelete(ctx context.Context, templateID string, opts *RequestOptions) error {
	_, err := r.client.do(ctx, "DELETE", "/v1/documents/templates/"+url.PathEscape(templateID), nil, nil, opts)
	return err
}

// Documents

func (r *DocumentsResource) Send(ctx context.Context, body interface{}, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/documents/send", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) SendToPartner(ctx context.Context, templateID, partnerID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/documents/send-to-partner", nil, map[string]string{"templateId": templateID, "partnerId": partnerID}, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) Get(ctx context.Context, documentID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents/"+url.PathEscape(documentID), nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) List(ctx context.Context, params url.Values, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents", params, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) GetStatus(ctx context.Context, documentID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents/"+url.PathEscape(documentID)+"/status", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) Download(ctx context.Context, documentID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents/"+url.PathEscape(documentID)+"/download", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) GetAuditTrail(ctx context.Context, documentID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "GET", "/v1/documents/"+url.PathEscape(documentID)+"/audit-trail", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) Void(ctx context.Context, documentID, reason string, opts *RequestOptions) (interface{}, error) {
	body := interface{}(nil)
	if reason != "" {
		body = map[string]string{"reason": reason}
	}
	resp, err := r.client.do(ctx, "POST", "/v1/documents/"+url.PathEscape(documentID)+"/void", nil, body, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *DocumentsResource) Resend(ctx context.Context, documentID string, opts *RequestOptions) (interface{}, error) {
	resp, err := r.client.do(ctx, "POST", "/v1/documents/"+url.PathEscape(documentID)+"/resend", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
