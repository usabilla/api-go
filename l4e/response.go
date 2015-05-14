package l4e

import (
	"encoding/json"

	"github.com/usabilla/gobilla/response"
)

// EmailButtonResponse is a response that contains email button data.
type EmailButtonResponse struct {
	response.Response
	Items []EmailButton `json:"items"`
}

// NewEmailButtonResponse creates an email button response and unmarshals
// json API email button response to Go struct.
func NewEmailButtonResponse(data []byte) (*EmailButtonResponse, error) {
	response := &EmailButtonResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
