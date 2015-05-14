package l4w

import (
	"encoding/json"

	"github.com/usabilla/gobilla/response"
)

// ButtonResponse is a response that contains button data.
type ButtonResponse struct {
	response.Response
	Items []Button `json:"items"`
}

// NewButtonResponse creates a button response and unmarshals json API
// button response to Go struct.
func NewButtonResponse(data []byte) (*ButtonResponse, error) {
	response := &ButtonResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// FeedbackResponse is a response that contains feedback item data.
type FeedbackResponse struct {
	response.Response
	Items []FeedbackItem `json:"items"`
}

// NewFeedbackResponse creates a feedback response and unmarshals json API
// feeddback items response to Go struct.
func NewFeedbackResponse(data []byte) (*FeedbackResponse, error) {
	response := &FeedbackResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CampaignResponse is a response that contains campaign data.
type CampaignResponse struct {
	response.Response
	Items []Campaign `json:"items"`
}

// NewCampaignResponse creates a campaign response and unmarshals json API
// campaign response to Go struct.
func NewCampaignResponse(data []byte) (*CampaignResponse, error) {
	response := &CampaignResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CampaignResultResponse is a response that contains campaign result data.
type CampaignResultResponse struct {
	response.Response
	Items []CampaignResult `json:"items"`
}

// NewCampaignResultResponse creates a new campaign result response and unmarshals json API
// campaign results response to Go struct.
func NewCampaignResultResponse(data []byte) (*CampaignResultResponse, error) {
	response := &CampaignResultResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CampaignStatsResponse is a response that contains campaign statistics data.
type CampaignStatsResponse struct {
	response.Response
	Items []CampaignStat `json:"items"`
}

// NewCampaignStatsResponse creates a new campaign statistics response and unmarshals json API
// campaign statistics response to Go struct.
func NewCampaignStatsResponse(data []byte) (*CampaignStatsResponse, error) {
	response := &CampaignStatsResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}