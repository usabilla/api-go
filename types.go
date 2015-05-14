package gobilla

import "time"

// Button represents a button item.
type Button struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FeedbackItem represents a feedback item.
type FeedbackItem struct {
	ID        string            `json:"id"`
	UserAgent string            `json:"userAgent"`
	Comment   string            `json:"comment"`
	Location  string            `json:"location"`
	Date      time.Time         `json:"date"`
	Custom    map[string]string `json:"custom"`
	Email     string            `json:"email"`
	Image     string            `json:"image,omitempty"`
	Labels    []string          `json:"labels"`
	NPS       int               `json:"nps"`
	PublicURL string            `json:"publicUrl"`
	Rating    int               `json:"rating"`
	ButtonID  string            `json:"buttonId"`
	Tags      []string          `json:"tags"`
	URL       string            `json:"url"`
}

// Campaign represents a campaign item.
type Campaign struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	ButtonID    string    `json:"buttonId"`
	AnalyticsID string    `json:"analyticsId"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
}

// CampaignResult represents a campaign result item.
type CampaignResult struct {
	ID         string                 `json:"id"`
	UserAgent  string                 `json:"userAgent"`
	Location   string                 `json:"location"`
	Date       time.Time              `json:"date"`
	CampaignID string                 `json:"campaignId"`
	Custom     map[string]string      `json:"custom"`
	Data       map[string]interface{} `json:"data"`
	URL        string                 `json:"url"`
	Time       float64                `json:"time"`
}

// CampaignStat represents a campaign statistics item
type CampaignStat struct {
	ID         string `json:"id"`
	Completed  int    `json:"completed"`
	Conversion int    `json:"conversion"`
	Views      int    `json:"views"`
}

// App represents an app item.
type App struct {
	ID     string `json:"id"`
	Date   string `json:"date"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// AppFeedbackItem represents an app feedback item.
type AppFeedbackItem struct {
	ID           string                 `json:"id"`
	Date         string                 `json:"date"`
	Timestamp    string                 `json:"timestamp"`
	DeviceName   string                 `json:"deviceName"`
	Data         map[string]interface{} `json:"data"`
	Custom       map[string]string      `json:"custom"`
	AppID        string                 `json:"appId"`
	AppName      string                 `json:"appName"`
	AppVersion   string                 `json:"appVersion"`
	OsName       string                 `json:"osName"`
	OsVersion    string                 `json:"osVersion"`
	Location     string                 `json:"location"`
	GeoLocation  map[string]interface{} `json:"geolocation"`
	FreeMemory   int                    `json:"freeMemory"`
	TotalMemory  int                    `json:"totalMemory"`
	FreeStorage  int                    `json:"freeStorage"`
	TotalStorage int                    `json:"totalStorage"`
	Screenshot   string                 `json:"screenshot"`
	Screensize   string                 `json:"screensize"`
	Connection   string                 `json:"connection"`
	IPAddress    string                 `json:"ipAddress"`
	Language     string                 `json:"language"`
	Orientation  string                 `json:"orientation"`
	BatteryLevel float32                `json:"batteryLevel"`
}

// EmailButton represents an email button item.
type EmailButton struct {
	ID        string                   `json:"id"`
	Date      string                   `json:"date"`
	Name      string                   `json:"name"`
	IntroText string                   `json:"introText"`
	Locale    string                   `json:"locale"`
	Groups    []map[string]interface{} `json:"groups"`
}
