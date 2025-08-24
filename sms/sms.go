package sms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type SMSSender struct {
	conf Config
}

type Request struct {
	To      string
	From    string
	Message string
}

type Response struct {
	ID     int    `json:"id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

func (s *SMSSender) SendSMS(ctx context.Context, req *Request) (*Response, error) {
	reqBody := map[string]string{
		"to":        req.To,
		"from":      req.From,
		"text":      req.Message,
		"apiKey":    s.conf.APIKey,
		"apiSecret": s.conf.APISecret,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.conf.Endpoint+"/sms", bytes.NewReader(jsonReqBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to send SMS")
	}
	var resp Response
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type Config struct {
	Endpoint  string
	APIKey    string
	APISecret string
}

func NewSMSSender(conf Config) *SMSSender {
	return &SMSSender{
		conf: conf,
	}
}
