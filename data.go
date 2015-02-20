package gobilla

/*
ButtonData ...
*/
type ButtonData struct {
	Items         []Button `json:"items"`
	Count         int      `json:"count"`
	HasMore       bool     `json:"hasMore"`
	LastTimestamp int64    `json:"lastTimestamp"`
}

/*
FeedbackData ...
*/
type FeedbackData struct {
	Items         []Feedback `json:"items"`
	Count         int        `json:"count"`
	HasMore       bool       `json:"hasMore"`
	LastTimestamp int64      `json:"lastTimestamp"`
}

/*
CampaignData ...
*/
type CampaignData struct {
	Items         []Campaign `json:"items"`
	Count         int        `json:"count"`
	HasMore       bool       `json:"hasMore"`
	LastTimestamp int64      `json:"lastTimestamp"`
}

/*
CampaignResultData ...
*/
type CampaignResultData struct {
	Items         []CampaignResult `json:"items"`
	Count         int              `json:"count"`
	HasMore       bool             `json:"hasMore"`
	LastTimestamp int64            `json:"lastTimestamp"`
}
