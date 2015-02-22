package gobilla

import "encoding/json"

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
	Response
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
	Response
	Items []FeedbackItem `json:"items"`
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
	Response
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
	Response
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
