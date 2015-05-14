package response

// Response contains common data for an API response.
type Response struct {
	Count         int   `json:"count"`
	HasMore       bool  `json:"hasMore"`
	LastTimestamp int64 `json:"lastTimestamp"`
}
