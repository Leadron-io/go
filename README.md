# Leadron Go SDK

Official Go SDK for the [Leadron](https://leadron.io) API — lead management and partner commission platform.

## Install

```bash
go get github.com/leadron/sdk-go
```

## Usage

```go
package main

import (
	"context"
	"github.com/leadron/sdk-go"
)

func main() {
	client := leadron.NewClient("your-api-key")
	// Or with config:
	// client := leadron.NewClientWithConfig(leadron.Config{
	//     APIKey:     "your-api-key",
	//     BaseURL:    "https://api.leadron.io",
	//     MaxRetries: 3,
	// })

	ctx := context.Background()

	// Auth
	valid, _ := client.Auth.Validate(ctx, nil)
	scopes, _ := client.Auth.GetScopes(ctx, nil)

	// Leads
	lead, _ := client.Leads.Create(ctx, map[string]interface{}{
		"email":     "jane@example.com",
		"firstName": "Jane",
		"lastName":  "Doe",
	}, nil)
	data, pagination, _ := client.Leads.List(ctx, &leadron.ListParams{Status: "qualified", Limit: 20}, nil)

	// Partners & commissions
	partner, _ := client.Partners.Get(ctx, partnerID, nil)
	client.Commissions.Approve(ctx, commissionID, nil)

	// Webhook signature verification (client-side)
	ok := leadron.VerifyWebhookSignature(rawBody, signature, secret)
	event, err := leadron.ConstructEvent(rawBody, signature, secret)

	// Optional: request ID and idempotency
	opts := &leadron.RequestOptions{RequestID: "my-request-id", IdempotencyKey: "unique-key-123"}
	client.Leads.Create(ctx, data, opts)

	// Rate limit (from last response)
	remaining := client.GetRateLimitStatus()
}
```

## API surface

- **Auth** — Validate, GetScopes; package-level VerifyWebhookSignature
- **Leads** — Create, Get, Update, Delete, List (returns data + pagination), Assign, UpdateStatus, AddNote, GetNotes, GetTimeline, MarkConverted, BulkCreate, BulkAssign, BulkUpdateStatus, Search, Filter
- **Partners** — Create, Get, Update, List, Deactivate, GetReferralTree, GetUpline, GetReferralLink, GetStats, GetLeaderboard, GetTopPerformers, Invite, ResendInvite, GetOnboardingStatus, SendAgreement, GetSignedDocuments, GetAgreementStatus
- **Commissions** — Create, Get, List, Approve, Reject, MarkPaid, GetRules, CreateRule, UpdateRule, DeleteRule, GetPayoutSummary, GetWalletBalance, RequestPayout, GetPayoutHistory, GetSummary, GetByPartner, GetTotalOwed, GetTotalPaid
- **Sequences** — Create, Get, List, Update, Delete, Activate, Pause, EnrollLead, EnrollBulk, UnenrollLead, GetEnrolledLeads, AddStep, UpdateStep, DeleteStep, ReorderSteps
- **SMS** — Send, GetInbox, GetOutbox, GetConversation, GetUsage
- **PhoneNumbers** — Search, List, Get, Release, AssignToTeam, UnassignFromTeam, GetUsage, Get10DLCStatus
- **Teams** — Create, Get, List, Update, Delete, AddMember, RemoveMember, GetMembers, AssignLead, AssignPhoneNumber, GetStats, GetLeaderboard
- **Documents** — TemplatesCreate, TemplatesList, TemplatesGet, TemplatesUpdate, TemplatesDelete, Send, SendToPartner, Get, List, GetStatus, Download, GetAuditTrail, Void, Resend
- **Webhooks** — Create, List, Get, Update, Delete, Test, GetLogs, Retry; package-level ConstructEvent
- **Analytics** — GetOverview, GetLeadMetrics, GetCommissionMetrics, GetPartnerMetrics, GetConversionRate, GetSmsMetrics
- **Reports** — Leads, Commissions, Partners, Export
- **Account** — Get, Update, GetBranding, UpdateBranding, ApiKeysList, ApiKeysCreate, ApiKeysRevoke, GetUsage, GetPlan, GetLimits

All methods accept `context.Context` as the first argument and `*RequestOptions` (RequestID, IdempotencyKey) as the last. Errors are *leadron.Error, *leadron.AuthError, *leadron.ValidationError, or *leadron.RateLimitError with StatusCode and Body.

## Webhook events

See [instructions.md](../instructions.md) for the full list of webhook event types. Verify payloads with `leadron.VerifyWebhookSignature(payload, signature, secret)` before processing; use `leadron.ConstructEvent` to verify and parse JSON.

## API docs

Full API reference: [OpenAPI spec](../../documentation/api-docs/openapi.yaml) and [SDK-to-HTTP mapping](../api-mapping.md).
