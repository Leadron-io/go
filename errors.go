package leadron

import "fmt"

// Error represents an API error with status code and body.
type Error struct {
	Message    string
	StatusCode int
	Body       interface{}
}

func (e *Error) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("%s (HTTP %d)", e.Message, e.StatusCode)
	}
	return e.Message
}

// AuthError is returned on 401.
type AuthError struct{ *Error }

// ValidationError is returned on 422.
type ValidationError struct{ *Error }

// RateLimitError is returned on 429 and includes RetryAfter.
type RateLimitError struct {
	*Error
	RetryAfter int
}
