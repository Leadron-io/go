package leadron

import (
	"net/http"
	"strings"
	"time"
)

// Leadron is the main API client. All resources are exposed as structs with methods.
type Leadron struct {
	Auth         *AuthResource
	Leads        *LeadsResource
	Partners     *PartnersResource
	Marketers    *MarketersResource
	Commissions  *CommissionsResource
	Sequences    *SequencesResource
	SMS          *SMSResource
	PhoneNumbers *PhoneNumbersResource
	Teams        *TeamsResource
	Documents    *DocumentsResource
	Webhooks     *WebhooksResource
	Analytics    *AnalyticsResource
	Reports      *ReportsResource
	Account      *AccountResource

	client *Client
}

// Config configures the Leadron client.
type Config struct {
	APIKey     string
	BaseURL    string
	MaxRetries int
}

// NewClient creates a new Leadron client. BaseURL defaults to https://api.leadron.io.
func NewClient(apiKey string) *Leadron {
	return NewClientWithConfig(Config{APIKey: apiKey})
}

// NewClientWithConfig creates a new Leadron client with full config.
func NewClientWithConfig(cfg Config) *Leadron {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	baseURL = strings.TrimSuffix(baseURL, "/")
	maxRetries := cfg.MaxRetries
	if maxRetries <= 0 {
		maxRetries = defaultRetries
	}
	rateLimit := new(int)
	c := &Client{
		apiKey:     cfg.APIKey,
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		maxRetries: maxRetries,
		rateLimit:  rateLimit,
	}
	return &Leadron{
		client:      c,
		Auth:        &AuthResource{client: c},
		Leads:       &LeadsResource{client: c},
		Partners:    &PartnersResource{client: c},
		Marketers:   &MarketersResource{client: c},
		Commissions: &CommissionsResource{client: c},
		Sequences:   &SequencesResource{client: c},
		SMS:         &SMSResource{client: c},
		PhoneNumbers: &PhoneNumbersResource{client: c},
		Teams:       &TeamsResource{client: c},
		Documents:   &DocumentsResource{client: c},
		Webhooks:    &WebhooksResource{client: c},
		Analytics:   &AnalyticsResource{client: c},
		Reports:     &ReportsResource{client: c},
		Account:     &AccountResource{client: c},
	}
}

// GetRateLimitStatus returns the last X-RateLimit-Remaining value from the API, if any.
func (l *Leadron) GetRateLimitStatus() *int {
	return l.client.GetRateLimitStatus()
}
