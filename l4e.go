/*
Copyright (c) 2015 Usabilla

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish, dis-
tribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the fol-
lowing conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABIL-
ITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT
SHALL THE AUTHOR BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
*/

package gobilla

import (
	"encoding/json"
	"net/http"
)

// Canonical URI constants.
const (
	emailURI       = "/live/email"
	emailButtonURI = emailURI + "/button"
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
	response
	Items []EmailButton `json:"items"`
}

// EmailButtons represents the email button resource of Usabilla API.
type EmailButtons struct {
	resource
	client http.Client
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

// Get function of EmailButtons resource returns all the email buttons
// taking into account the specified query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (eb *EmailButtons) Get(params map[string]string) (*EmailButtonResponse, error) {
	request := request{
		method: "GET",
		auth:   eb.auth,
		uri:    emailButtonURI,
		params: params,
		client: eb.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewEmailButtonResponse(data)
}

// Feedback encapsulates the email feedback item resource.
//
// We use the FeedbackItems subresource of websites button feedback as it is
// the same. The only difference being that image is contained
// in the website feedback item response, but it is omitted for the email one.
func (eb *EmailButtons) Feedback() *FeedbackItems {
	return &FeedbackItems{
		resource: resource{
			auth: eb.auth,
		},
	}
}
