package model_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/sean-rn/hellosign-sdk/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventCallbackRequestEvent(t *testing.T) {
	jsonBytes, err := os.ReadFile("testdata/signature_request_all_signed.json")
	require.NoError(t, err)

	var actual model.EventCallbackRequest
	err = json.Unmarshal(jsonBytes, &actual)
	require.NoError(t, err)

	expectedCreatedAt := time.Date(2024, time.October, 28, 18, 26, 37, 0, time.UTC).In(time.Local)
	assert.Equal(t, expectedCreatedAt, actual.Event.EventTime.Time)
}
