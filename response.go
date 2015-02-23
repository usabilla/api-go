package gobilla

import "encoding/json"

// Response contains common data for an API response.
type Response struct {
	Count         int   `json:"count"`
	HasMore       bool  `json:"hasMore"`
	LastTimestamp int64 `json:"lastTimestamp"`
}

// ButtonResponse is a response that contains button data.
type ButtonResponse struct {
	Response
	Items []Button `json:"items"`
}

// JSON response.
func (response *ButtonResponse) JSON(data []byte) (*ButtonResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FeedbackResponse is a response that contains feedback item data.
type FeedbackResponse struct {
	Response
	Items []FeedbackItem `json:"items"`
}

// JSON response.
func (response *FeedbackResponse) JSON(data []byte) (*FeedbackResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// CampaignResponse is a response that contains campaign data.
type CampaignResponse struct {
	Response
	Items []Campaign `json:"items"`
}

// JSON response.
func (response *CampaignResponse) JSON(data []byte) (*CampaignResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// CampaignResultResponse is a response that contains campaign result data.
type CampaignResultResponse struct {
	Response
	Items []CampaignResult `json:"items"`
}

// JSON response.
func (response *CampaignResultResponse) JSON(data []byte) (*CampaignResultResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
