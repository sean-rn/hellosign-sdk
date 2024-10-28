package model

import (
	"encoding/json"
)

// EventCallbackRequest struct for EventCallbackRequest
type EventCallbackRequest struct {
	// Basic information about the event that occurred.
	Event EventCallbackRequestEvent `json:"event"`
	// Contains information about a signature request.
	SignatureRequest *SignatureRequestResponse `json:"signature_request,omitempty"`
	// Contains information about the accounts you and your team have created. (NOT IMPLEMENTED)
	Account json.RawMessage `json:"account,omitempty"`
	// Contains information about the templates you and your team have created. (NOT IMPLEMENTED)
	Template json.RawMessage `json:"template,omitempty"`
}
