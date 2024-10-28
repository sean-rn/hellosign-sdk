package model

// EventCallbackRequestEvent Basic information about the event that occurred.
type EventCallbackRequestEvent struct {
	// Time the event was created (using Unix time).
	EventTime UnixTimestamp `json:"event_time"`
	// Type of callback event that was triggered.
	EventType string `json:"event_type"`
	// Generated hash used to verify source of event data.
	EventHash string `json:"event_hash"`
	// Specific metadata about the event.
	EventMetadata *EventCallbackRequestEventMetadata `json:"event_metadata,omitempty"`
}

// EventCallbackRequestEventMetadata Specific metadata about the event.
type EventCallbackRequestEventMetadata struct {
	// Signature ID for a specific signer. Applicable to `signature_request_signed` and `signature_request_viewed` events.
	RelatedSignatureId string `json:"related_signature_id,omitempty"`
	// Account ID the event was reported for.
	ReportedForAccountId string `json:"reported_for_account_id,omitempty"`
	// App ID the event was reported for.
	ReportedForAppId string `json:"reported_for_app_id,omitempty"`
	// Message about a declined or failed (due to error) signature flow.
	EventMessage string `json:"event_message,omitempty"`
}
