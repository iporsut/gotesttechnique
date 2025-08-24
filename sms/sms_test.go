package sms

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMSSender_SendSMS_Success(t *testing.T) {
	smsSender := NewSMSSender(Config{
		Endpoint:  "https://api.example.com",
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
