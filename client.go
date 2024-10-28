package hellosign

import (
	"context"
	"net/http"

	"github.com/sean-rn/hellosign-sdk/model"
)

const (
	DefaultBaseURL = "https://api.hellosign.com/v3"
)

// API declares the methods of Client as an interface for your convenience.
type API interface {
	// DownloadFiles Obtain a copy of the current documents specified by the `signature_request_id` parameter.
	// Returns a PDF or ZIP file. If the files are currently being prepared, a status code of `409` will be returned instead.
	// Parameters:
	//   - signatureRequestId The id of the SignatureRequest to retrieve.
	//   - fileType Set to "pdf" for a single merged document or "zip" for a collection of individual documents.
	DownloadFiles(ctx context.Context, signatureRequestId, fileType string) ([]byte, error)

	// Create Embedded Signature Request with Template
	// Creates a new SignatureRequest based on the given Template(s) to be signed in an embedded iFrame.
	// Note that embedded signature requests can only be signed in embedded iFrames whereas normal signature requests
	// can only be signed on Dropbox Sign.
	CreateEmbeddedWithTemplate(ctx context.Context, req model.CreateEmbeddedWithTemplateRequest) (*model.SignatureRequestGetResponse, error)

	// Retrieves an embedded object containing a signature url that can be opened in an iFrame.
	// Parameters:
	//   - signatureId The id of the signature to get a signature url for.
	GetEmbeddedSignUrl(ctx context.Context, signatureId string) (*model.EmbeddedSignUrlResponse, error)
}

// Assert that *Client implements API
var _ API = (*Client)(nil)

type Client struct {
	httpClient *http.Client                  // A custom *http.Client to use, otherwise use http.DefaultClient
	signer     func(req *http.Request) error // signer adds authentication header(s) to the request, returning an error if it can't
	baseURL    string                        // Base URL to which to append endpoint paths
}

// NewClient creates a new Hellosign API client with optional configuration options.
func NewClient(options ...Option) *Client {
	c := new(Client)
	for _, option := range options {
		option(c)
	}
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	if c.baseURL == "" {
		c.baseURL = DefaultBaseURL
	}
	return c
}

type Option func(*Client)

// WithHTTPClient configures the client to use hc instead of http.DefaultClient
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithApiKey configures the client to authenticate using a Hellosign API key, which can be
// which can be retrieved from the [API tab] of your [API Settings page].
//
// [API Settings page]: https://app.hellosign.com/home/myAccount#api
func WithApiKey(key string) Option {
	return func(c *Client) {
		c.signer = func(req *http.Request) error {
			req.SetBasicAuth(key, "")
			return nil
		}
	}
}

// WithAccessToken configures the client to authenticate using an access token (issued during
// an OAuth flow) to send API requests on behalf of the user that granted authorization.
func WithAccessToken(token string) Option {
	return func(c *Client) {
		c.signer = func(req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			return nil
		}
	}
}

// WithBaseURL uses baseURL as the baseURL instead of [DefaultBaseURL].
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
