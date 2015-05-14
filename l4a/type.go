package l4a

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
