/*
Copyright (c) 2018 Usabilla

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

package usabilla

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Canonical URI constants.
const (
	appsURI = "/live/apps"
)

var (
	appFeedbackURI = appsURI + "/%s/feedback"
)

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

// Apps represents the app resource of Usabilla API.
type Apps struct {
	resource
	client *http.Client
}

// AppFeedbackItems represents the apps feedback item subresource of Usabilla API.
type AppFeedbackItems struct {
	resource
	client *http.Client
}

// AppResponse is a response that contains app data.
type AppResponse struct {
	response
	Items []App `json:"items"`
}

// AppFeedbackResponse is a response that contains app feedback item data.
type AppFeedbackResponse struct {
	response
	Items []AppFeedbackItem `json:"items"`
}

// Get function of Apps resource returns all apps
// taking into account the specified query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (a *Apps) Get(params map[string]string) (*AppResponse, error) {
	request := request{
		method: "GET",
		auth:   a.auth,
		uri:    appsURI,
		params: params,
		client: a.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return newAppResponse(data)
}

// Feedback encapsulates the app feedback item resource.
func (a *Apps) Feedback() *AppFeedbackItems {
	return &AppFeedbackItems{
		resource: resource{
			auth: a.auth,
		},
		client: a.client,
	}
}

// Get function of AppFeedbackItem resource returns all the feedback items
// for a specific app, taking into account the provided query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (af *AppFeedbackItems) Get(appID string, params map[string]string) (*AppFeedbackResponse, error) {
	uri := fmt.Sprintf(appFeedbackURI, appID)

	request := &request{
		method: "GET",
		auth:   af.auth,
		uri:    uri,
		params: params,
		client: af.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return newAppFeedbackResponse(data)
}

// Iterate uses an AppFeedbackItem channel which transparently uses the HasMore field to fire
// a new api request once all items have been consumed on the channel.
func (af *AppFeedbackItems) Iterate(appID string, params map[string]string) chan AppFeedbackItem {
	resp, err := af.Get(appID, params)

	if err != nil {
		panic(err)
	}

	afic := make(chan AppFeedbackItem)

	go appItems(afic, resp, af, appID)

	return afic
}

// appItems feeds a feedback item channel with items.
//
// While hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items.
//
// When HasMore is false, we close the channel.
func appItems(afic chan AppFeedbackItem, resp *AppFeedbackResponse, af *AppFeedbackItems, appID string) {
	for {
		for _, item := range resp.Items {
			afic <- item
		}
		if !resp.HasMore {
			close(afic)
			return
		}
		params := map[string]string{
			"since": strconv.FormatInt(resp.LastTimestamp, 10),
		}

		resp, err := af.Get(appID, params)

		if err != nil {
			panic(err)
		}

		appItems(afic, resp, af, appID)

		return
	}
}

// NewAppResponse creates an app response and unmarshals json API app
// response to Go struct.
func newAppResponse(data []byte) (*AppResponse, error) {
	response := &AppResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewAppFeedbackResponse creates an app feedback response and unmarshals json
// API app feeddback items response to Go struct.
func newAppFeedbackResponse(data []byte) (*AppFeedbackResponse, error) {
	response := &AppFeedbackResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
