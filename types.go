package gobilla

import "time"

// Button represents a button item.
type Button struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FeedbackItem represents a feedback item.
type FeedbackItem struct {
	ID          string            `json:"id"`
	UserAgent   string            `json:"userAgent"`
	Comment     string            `json:"comment"`
	Location    string            `json:"location"`
	Date        time.Time         `json:"date"`
	Custom      map[string]string `json:"custom"`
	Email       string            `json:"email"`
	HTMLSnippet string            `json:"htmlSnippet"`
	Image       string            `json:"image"`
	Labels      []string          `json:"labels"`
	NPS         int               `json:"nps"`
	PublicURL   string            `json:"publicUrl"`
	Rating      int               `json:"rating"`
	ButtonID    string            `json:"buttonId"`
	Tags        []string          `json:"tags"`
	URL         string            `json:"url"`
}

// Campaign represents a campaign item.
type Campaign struct {
	ID             string    `json:"id"`
	Date           time.Time `json:"date"`
	FeedbackSiteID string    `json:"buttonId"`
	JsID           string    `json:"analyticsId"`
	Status         string    `json:"status"`
	Title          string    `json:"name"`
}

// CampaignResult represents a campaign result item.
type CampaignResult struct {
	ID         string                 `json:"id"`
	UserAgent  string                 `json:"userAgent"`
	Location   string                 `json:"location"`
	Date       time.Time              `json:"date"`
	CampaignID string                 `json:"campaignId"`
	Custom     map[string]string      `json:"customData"`
	Data       map[string]interface{} `json:"data"`
	URL        string                 `json:"url"`
	TotalTime  float64                `json:"time"`
}
