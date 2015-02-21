package gobilla

import "encoding/json"

/*
BaseData ...
*/
type BaseData struct {
	Count         int   `json:"count"`
	HasMore       bool  `json:"hasMore"`
	LastTimestamp int64 `json:"lastTimestamp"`
}

/*
Data ...
*/
type Data interface {
	JSON(response []byte) (Data, error)
}

/*
ButtonData ...
*/
type ButtonData struct {
	BaseData
	Items []Button `json:"items"`
}

/*
JSON ...
*/
func (data ButtonData) JSON(response []byte) (ButtonData, error) {
	err := json.Unmarshal(response, &data)
	if err != nil {
		return ButtonData(data), err
	}
	return ButtonData(data), nil
}

/*
FeedbackData ...
*/
type FeedbackData struct {
	BaseData
	Items []Feedback `json:"items"`
}

/*
JSON ...
*/
func (data FeedbackData) JSON(response []byte) (FeedbackData, error) {
	err := json.Unmarshal(response, &data)
	if err != nil {
		return FeedbackData(data), err
	}
	return FeedbackData(data), nil
}

/*
CampaignData ...
*/
type CampaignData struct {
	BaseData
	Items []Campaign `json:"items"`
}

/*
JSON ...
*/
func (data CampaignData) JSON(response []byte) (CampaignData, error) {
	err := json.Unmarshal(response, &data)
	if err != nil {
		return CampaignData(data), err
	}
	return CampaignData(data), nil
}

/*
CampaignResultData ...
*/
type CampaignResultData struct {
	BaseData
	Items []CampaignResult `json:"items"`
}

/*
JSON ...
*/
func (data CampaignResultData) JSON(response []byte) (CampaignResultData, error) {
	err := json.Unmarshal(response, &data)
	if err != nil {
		return CampaignResultData(data), err
	}
	return CampaignResultData(data), nil
}
