package gobilla

import "time"

/*
Button represents a button item.
*/
type Button struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/*
FeedbackItem represents a feedback item.
*/
type FeedbackItem struct {
	ID        string            `json:"id"`
	UserAgent string            `json:"userAgent"`
	Comment   string            `json:"comment"`
	Location  string            `json:"location"`
	Date      time.Time         `json:"date"`
	Custom    map[string]string `json:"custom"`
	Email     string            `json:"email"`
	Image     string            `json:"image"`
	Labels    []string          `json:"labels" json:"labels"`
	NPS       string            `json:"nps" json:"nps"`
	PublicURL string            `json:"public_id" json:"publicUrl"`
	Rating    string            `json:"rating" json:"rating"`
	ButtonID  string            `json:"site_id" json:"buttonId"`
	Tags      []string          `json:"tags" json:"tags"`
	URL       string            `json:"url" json:"url"`
}

/*
Campaign represents a campaign item.
*/
type Campaign struct {
	ID             string    `json:"id"`
	Date           time.Time `json:"date"`
	FeedbackSiteID string    `json:"buttonId"`
	JsID           string    `json:"analyticsId"`
	Status         string    `json:"status"`
	Title          string    `json:"name"`
}

/*
CampaignResult represents a campaign result item.
*/
type CampaignResult struct {
	ID         string            `json:"id"`
	UserAgent  string            `json:"userAgent"`
	Location   string            `json:"location"`
	Date       time.Time         `json:"date"`
	CampaignID string            `json:"campaignId"`
	Custom     map[string]string `json:"customData"`
	Data       map[string]string `json:"data"`
	URL        string            `json:"url"`
	TotalTime  int64             `json:"time"`
}

// Response contains common data for an API response.
type Response struct {
	Count         int   `json:"count"`
	HasMore       bool  `json:"hasMore"`
	LastTimestamp int64 `json:"lastTimestamp"`
}
