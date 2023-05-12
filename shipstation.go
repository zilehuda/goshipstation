package shipstation

import (
	"bytes"
	"net/http"
)

// ShipStation represents the main client for interacting with ShipStation API.
type ShipStation struct {
	apiKey    string
	apiSecret string
	baseURL   string
}

// NewShipStation creates a new ShipStation client with the specified API credentials.
func NewShipStation(apiKey, apiSecret string) *ShipStation {
	return &ShipStation{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   "https://ssapi.shipstation.com",
	}
}

// SetBaseURL sets the base URL for ShipStation API.
func (s *ShipStation) SetBaseURL(baseURL string) {
	s.baseURL = baseURL
}

func (s *ShipStation) sendRequest(method, urlStr string, actionName string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	s.setBasicAuth(req)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ShipStation) setBasicAuth(req *http.Request) {
	req.SetBasicAuth(s.apiKey, s.apiSecret)
}
