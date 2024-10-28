package hellosign_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sean-rn/hellosign-sdk"
	"github.com/sean-rn/hellosign-sdk/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	// Create mock API Server
	server := setupMockAPIServer()
	t.Cleanup(server.Close)

	ctx := context.Background()
	client := hellosign.NewClient(hellosign.WithBaseURL(server.URL), hellosign.WithApiKey("test-api-key"))
	srResp, err := client.CreateEmbeddedWithTemplate(ctx, model.CreateEmbeddedWithTemplateRequest{
		ClientId:    "ddddb5e5c34b929957e24b17aa52dddd",
		TemplateIds: []string{"cccc6ad681229567aab20cd83a69cf18fb2cccc"},
		Signers: []model.SubSignatureRequestTemplateSigner{
			{Role: "First", Name: "Signer One", EmailAddress: "signer.one@example.org"},
		},
		Metadata: map[string]interface{}{
			"partner_user_id": 3456,
		},
		TestMode: true,
	})
	require.NoError(t, err)
	require.NotNil(t, srResp)

	assert.Equal(t, "ebaae602348695a4c712aa0f22614986d03caaaa", srResp.SignatureRequest.SignatureRequestId)
	expectedCreatedAt := time.Date(2024, time.October, 28, 17, 40, 43, 0, time.UTC).In(time.Local)
	if assert.NotNil(t, srResp.SignatureRequest.CreatedAt) {
		assert.Equal(t, expectedCreatedAt, srResp.SignatureRequest.CreatedAt.Time)
	}
}

func setupMockAPIServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v3/signature_request/create_embedded_with_template", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/create_embedded_with_template.resp.json")
	})
	return httptest.NewServer(mux)
}
