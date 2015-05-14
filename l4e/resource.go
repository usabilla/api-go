package l4e

import (
	"github.com/usabilla/gobilla/l4w"
	"github.com/usabilla/gobilla/request"
	"github.com/usabilla/gobilla/resource"
)

// Canonical URI constants.
const (
	emailURI       = "/live/email"
	emailButtonURI = emailURI + "/button"
)

// EmailButtons represents the email button resource of Usabilla API.
type EmailButtons struct {
	resource.Resource
}

// Get function of EmailButtons resource returns all the email buttons
// taking into account the specified query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (eb *EmailButtons) Get(params map[string]string) (*EmailButtonResponse, error) {
	request := request.Request{
		Method: "GET",
		Auth:   eb.Auth,
		URI:    emailButtonURI,
		Params: params,
	}

	data, err := request.Get()
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
func (eb *EmailButtons) Feedback() *l4w.FeedbackItems {
	return &l4w.FeedbackItems{
		Resource: resource.Resource{
			Auth: eb.Auth,
		},
	}
}
