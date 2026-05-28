package leadron

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// AuthResource handles auth and webhook verification.
type AuthResource struct {
	client *Client
}

// Validate validates the API key. Returns valid true if the key is valid.
func (r *AuthResource) Validate(ctx context.Context, opts *RequestOptions) (valid bool, err error) {
	resp, err := r.client.do(ctx, "GET", "/v1/api-keys/validate", nil, nil, opts)
	if err != nil {
		return false, err
	}
	if m, ok := resp.Data.(map[string]interface{}); ok && m["valid"] != nil {
		if v, ok := m["valid"].(bool); ok {
			return v, nil
		}
	}
	return true, nil
}

// GetScopes returns the scopes for the current API key.
func (r *AuthResource) GetScopes(ctx context.Context, opts *RequestOptions) (scopes []string, err error) {
	resp, err := r.client.do(ctx, "GET", "/v1/api-keys/scopes", nil, nil, opts)
	if err != nil {
		return nil, err
	}
	if m, ok := resp.Data.(map[string]interface{}); ok && m["scopes"] != nil {
		if s, ok := m["scopes"].([]interface{}); ok {
			for _, v := range s {
				if str, ok := v.(string); ok {
					scopes = append(scopes, str)
				}
			}
		}
	}
	return scopes, nil
}

// VerifyWebhookSignature verifies an HMAC-SHA256 signature over the raw body.
func VerifyWebhookSignature(payload, signature, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(signature)) ||
		hmac.Equal([]byte("sha256="+expected), []byte(signature))
}
