package hellosign

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sean-rn/hellosign-sdk/model"
)

// DownloadFiles Obtain a copy of the current documents specified by the `signature_request_id` parameter.
// Returns a PDF or ZIP file. If the files are currently being prepared, a status code of `409` will be returned instead.
// Parameters:
//   - signatureRequestId The id of the SignatureRequest to retrieve.
//   - fileType Set to "pdf" for a single merged document or "zip" for a collection of individual documents.
func (c *Client) DownloadFiles(ctx context.Context, signatureRequestId, fileType string) ([]byte, error) {
	furl := fmt.Sprintf("%s/v3/signature_request/files/%s", c.baseURL, url.PathEscape(signatureRequestId))
	if fileType != "" {
		furl += "?file_type=" + url.QueryEscape(fileType)
	}

	req, err := c.newJSONRequest(ctx, http.MethodGet, furl, nil)
	if err != nil {
		return nil, err
	}
	var data []byte
	err = c.doRequest(req, &data)
	return data, err
}

// Create Embedded Signature Request with Template
// Creates a new SignatureRequest based on the given Template(s) to be signed in an embedded iFrame.
// Note that embedded signature requests can only be signed in embedded iFrames whereas normal signature requests
// can only be signed on Dropbox Sign.
func (c *Client) CreateEmbeddedWithTemplate(ctx context.Context, r model.CreateEmbeddedWithTemplateRequest) (*model.SignatureRequestGetResponse, error) {
	furl := fmt.Sprintf("%s/v3/signature_request/create_embedded_with_template", c.baseURL)
	req, err := c.newJSONRequest(ctx, http.MethodPost, furl, r)
	if err != nil {
		return nil, err
	}
	var resp model.SignatureRequestGetResponse
	err = c.doRequest(req, &resp)
	return &resp, err
}

// Retrieves an embedded object containing a signature url that can be opened in an iFrame.
// Parameters:
//   - signatureId The id of the signature to get a signature url for.
func (c *Client) GetEmbeddedSignUrl(ctx context.Context, signatureId string) (*model.EmbeddedSignUrlResponse, error) {
	furl := fmt.Sprintf("%s/v3/embedded/sign_url/%s", c.baseURL, url.PathEscape(signatureId))
	req, err := c.newJSONRequest(ctx, http.MethodPost, furl, nil)
	if err != nil {
		return nil, err
	}
	var resp model.EmbeddedSignUrlResponse
	err = c.doRequest(req, &resp)
	return &resp, err
}

// newJSONRequest creates a signed request with an optional JSON request body
func (c *Client) newJSONRequest(ctx context.Context, method, url string, body any) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonStr, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshalling body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonStr)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	if c.signer != nil {
		if err := c.signer(req); err != nil {
			return nil, fmt.Errorf("signing request: %w", err)
		}
	}
	return req, nil
}

// Do sends an HTTP request and optionally parses the response into a target.
func (c *Client) doRequest(req *http.Request, target any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, err := io.ReadAll(io.LimitReader(resp.Body, 4096))
		if err != nil {
			return fmt.Errorf("error reading error response: %s: %w", resp.Status, err)
		}
		return fmt.Errorf("request returned %s: %s", resp.Status, respBody)
	}

	switch t := target.(type) {
	case nil:
		return nil // Do nothing, no target given
	case *[]byte:
		var err error
		*t, err = io.ReadAll(resp.Body)
		return err
	default:
		return json.NewDecoder(resp.Body).Decode(target)
	}
}
