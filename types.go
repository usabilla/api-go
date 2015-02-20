package gobilla

import "time"

/*
Feedback ...
*/
type Feedback struct {
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
Campaign ...
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
CampaignResult ...
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

/*
Button ...
*/
type Button struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
