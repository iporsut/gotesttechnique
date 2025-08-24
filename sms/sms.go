package sms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Sender struct {
	endpoint   string
	apiKey     string
	apiSecret  string
	httpClient *http.Client
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

func (s *Sender) SendSMS(ctx context.Context, req *Request) (*Response, error) {
	reqBody := map[string]string{
		"to":        req.To,
		"from":      req.From,
		"text":      req.Message,
		"apiKey":    s.apiKey,
		"apiSecret": s.apiSecret,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint+"/sms", bytes.NewReader(jsonReqBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := s.httpClient.Do(httpReq)
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
	Endpoint   string
	APIKey     string
	APISecret  string
	HTTPClient *http.Client
}

func NewSender(conf Config) *Sender {
	if conf.HTTPClient == nil {
		conf.HTTPClient = http.DefaultClient
	}

	return &Sender{
		endpoint:   conf.Endpoint,
		apiKey:     conf.APIKey,
		apiSecret:  conf.APISecret,
		httpClient: conf.HTTPClient,
	}
}
