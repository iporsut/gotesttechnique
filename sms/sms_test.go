package sms

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func smsTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{
			"id": 1,
			"from": "TestSender",
			"to": "+1234567890",
			"text": "Hello, this is a test message.",
			"status": "ACCEPTED"
		}`))
	}))
}

func TestSMSSender_SendSMS_Success(t *testing.T) {
	server := smsTestServer()
	defer server.Close()

	smsSender := NewSMSSender(Config{
		Endpoint:  server.URL,
		APIKey:    "test-api",
		APISecret: "test-secret",
	})

	resp, err := smsSender.SendSMS(context.Background(), &Request{
		To:      "+1234567890",
		From:    "TestSender",
		Message: "Hello, this is a test message.",
	})

	expectedResp := &Response{
		ID:     1,
		From:   "TestSender",
		To:     "+1234567890",
		Text:   "Hello, this is a test message.",
		Status: "ACCEPTED",
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}
