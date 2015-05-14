package l4a

import (
	"encoding/json"

	"github.com/usabilla/gobilla/response"
)

// AppResponse is a response that contains app data.
type AppResponse struct {
	response.Response
	Items []App `json:"items"`
}

// NewAppResponse creates an app response and unmarshals json API app
// response to Go struct.
func NewAppResponse(data []byte) (*AppResponse, error) {
	response := &AppResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// AppFeedbackResponse is a response that contains app feedback item data.
type AppFeedbackResponse struct {
	response.Response
	Items []AppFeedbackItem `json:"items"`
}

// NewAppFeedbackResponse creates an app feedback response and unmarshals json
// API app feeddback items response to Go struct.
func NewAppFeedbackResponse(data []byte) (*AppFeedbackResponse, error) {
	response := &AppFeedbackResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
