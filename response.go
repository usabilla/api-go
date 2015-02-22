package gobilla

import "encoding/json"

// Response contains common data for an API response.
type Response struct {
	Count         int   `json:"count"`
	HasMore       bool  `json:"hasMore"`
	LastTimestamp int64 `json:"lastTimestamp"`
}

/*
ButtonResponse ...
*/
type ButtonResponse struct {
	Response
	Items []Button `json:"items"`
}

/*
JSON ...
*/
func (response *ButtonResponse) JSON(data []byte) (*ButtonResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

/*
FeedbackResponse ...
*/
type FeedbackResponse struct {
	Response
	Items []FeedbackItem `json:"items"`
}

/*
JSON ...
*/
func (response *FeedbackResponse) JSON(data []byte) (*FeedbackResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

/*
CampaignResponse ...
*/
type CampaignResponse struct {
	Response
	Items []Campaign `json:"items"`
}

/*
JSON ...
*/
func (response *CampaignResponse) JSON(data []byte) (*CampaignResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

/*
CampaignResultResponse ...
*/
type CampaignResultResponse struct {
	Response
	Items []CampaignResult `json:"items"`
}

/*
JSON ...
*/
func (response *CampaignResultResponse) JSON(data []byte) (*CampaignResultResponse, error) {
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
