package model

// EmbeddedSignUrlResponse struct for EmbeddedSignUrlResponse
type EmbeddedSignUrlResponse struct {
	Embedded EmbeddedSignUrlResponseEmbedded `json:"embedded"`
	Warnings []WarningResponse               `json:"warnings,omitempty"` // A list of warnings.
}

// EmbeddedSignUrlResponseEmbedded An object that contains necessary information to set up embedded signing.
type EmbeddedSignUrlResponseEmbedded struct {
	// A signature url that can be opened in an iFrame.
	SignURL string `json:"sign_url,omitempty"`
	// The specific time that the the `sign_url` link expires, in epoch.
	ExpiresAt *UnixTimestamp `json:"expires_at,omitempty"`
}
