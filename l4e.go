package gobilla

import (
	"encoding/json"
)

// EmailButton represents an email button item.
type EmailButton struct {
	ID        string                   `json:"id"`
	Date      string                   `json:"date"`
	Name      string                   `json:"name"`
	IntroText string                   `json:"introText"`
	Locale    string                   `json:"locale"`
	Groups    []map[string]interface{} `json:"groups"`
}

// EmailButtonResponse is a response that contains email button data.
type EmailButtonResponse struct {
	Response
	Items []EmailButton `json:"items"`
}

// NewEmailButtonResponse creates an email button response and unmarshals
// json API email button response to Go struct.
func NewEmailButtonResponse(data []byte) (*EmailButtonResponse, error) {
	response := &EmailButtonResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Canonical URI constants.
const (
	emailURI       = "/live/email"
	emailButtonURI = emailURI + "/button"
)

// EmailButtons represents the email button resource of Usabilla API.
type EmailButtons struct {
	resource
}

// Get function of EmailButtons resource returns all the email buttons
// taking into account the specified query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (eb *EmailButtons) Get(params map[string]string) (*EmailButtonResponse, error) {
	request := request{
		method: "GET",
		auth:   eb.auth,
		uri:    emailButtonURI,
		params: params,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewEmailButtonResponse(data)
}

// Feedback encapsulates the email feedback item resource.
//
// We use the FeedbackItem response as it is the same with the feedback item
// response from websites, only difference is that image is contained
// in the website feedback item response, but it is omitted for the email one
func (eb *EmailButtons) Feedback() *FeedbackItems {
	return &FeedbackItems{
		resource: resource{
			auth: eb.auth,
		},
	}
}
