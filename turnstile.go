package turnstile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Allow adjusting endpoint url
var SiteVerifyEndpoint = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

// A wrapper client around Cloudflare Turnstile
type TurnstileClient struct {
	*http.Client
}

// The response after a turnstile verify request
type VerifyResponse struct {
	Success     bool     `json:"success"`
	ErrorCodes  []string `json:"error-codes"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
}

// Available options for verify requests
type VerifyOpts struct {
	Secret         string `json:"secret"`          // required
	Response       string `json:"response"`        // required
	RemoteIP       string `json:"remoteip"`        // optional
	IdemPotencyKey string `json:"idempotency_key"` // optional
}

// Send verify requests using the given options
func (client *TurnstileClient) Verify(options VerifyOpts) (*VerifyResponse, error) {
	payload, err := json.Marshal(&options)
	if err != nil {
		return nil, err
	}

	payload_buff := bytes.NewBuffer(payload)
	req, err := http.NewRequest(http.MethodPost, SiteVerifyEndpoint, payload_buff)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("err: failed with status %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	results := VerifyResponse{}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

// Check response validity
func (client *TurnstileClient) Valid(results VerifyResponse) bool {
	return results.Success
}
