package l4e

// EmailButton represents an email button item.
type EmailButton struct {
	ID        string                   `json:"id"`
	Date      string                   `json:"date"`
	Name      string                   `json:"name"`
	IntroText string                   `json:"introText"`
	Locale    string                   `json:"locale"`
	Groups    []map[string]interface{} `json:"groups"`
}
