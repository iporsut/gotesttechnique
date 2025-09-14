package sms

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func smsTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request:", r.Method, r.URL.Path)
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
	server := smsTestServer() // สร้าง test server
	defer server.Close()

	smsSender := NewSender(Config{
		Endpoint:   server.URL, // test server endpoint
		APIKey:     "test-api",
		APISecret:  "test-secret",
		HTTPClient: server.Client(),
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

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, expectedResp, resp, "expected response to match")
}
