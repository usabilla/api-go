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
	appsURI         = "/live/apps"
	appsCampaignURI = appsURI + "/campaign"
)

var (
	appFeedbackURI        = appsURI + "/%s/feedback"
	appCampaignResultsURI = appsCampaignURI + "/%s/results"
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

// AppCampaignStruct represents a campaign item.
type AppCampaignStruct struct {
	ID             string   `json:"id"`
	CreatedAt      string   `json:"createdAt"`
	LastModifiedAt string   `json:"lastModifiedAt"`
	Status         string   `json:"status"`
	Name           string   `json:"name"`
	AppIds         []string `json:"appIds"`
}

// AppCampaignResultStruct represents a campaign result item.
type AppCampaignResultStruct struct {
	ID         string                 `json:"id"`
	Date       string                 `json:"date"`
	CampaignID string                 `json:"campaignId"`
	AppID      string                 `json:"appId"`
	Data       map[string]interface{} `json:"data"`
	Context    map[string]interface{} `json:"context"`
	Metadata   map[string]interface{} `json:"metadata"`
	Complete   bool                   `json:"complete"`
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

// AppCampaigns represents a App Campaign resource.
type AppCampaigns struct {
	resource
	client *http.Client
}

// AppCampaignResults represents a App Campaign Results resource.
type AppCampaignResults struct {
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

// AppCampaignResponse is a response that contains App campaign data.
type AppCampaignResponse struct {
	response
	Items []AppCampaignStruct `json:"items"`
}

// AppCampaignResults is a response that contains App campaign results data.
type AppCampaignResultsResponse struct {
	response
	Items []AppCampaignResultStruct `json:"items"`
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

// Get function of AppCampaigns resource returns all the campaigns
func (ac *AppCampaigns) Get(params map[string]string) (*AppCampaignResponse, error) {
	request := &request{
		method: "GET",
		auth:   ac.auth,
		uri:    appsCampaignURI,
		params: params,
		client: ac.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return newAppCampaignResponse(data)
}

// Results encapsulates the App Campaign Results resource
func (ac *AppCampaigns) Results() *AppCampaignResults {
	return &AppCampaignResults{
		resource: resource{
			auth: ac.auth,
		},
		client: ac.client,
	}
}

// Get function of AppCampaignResults resource returns all the campaign results
func (acr *AppCampaignResults) Get(campaignID string, params map[string]string) (*AppCampaignResultsResponse, error) {
	uri := fmt.Sprintf(appCampaignResultsURI, campaignID)

	request := &request{
		method: "GET",
		auth:   acr.auth,
		uri:    uri,
		params: params,
		client: acr.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return newAppCampaignResultsResponse(data)
}

// Iterate uses a channel which transparently uses the HasMore field to fire a new API request. Once all results
// have consumed on the channel it closes the chanel.
func (ac *AppCampaignResults) Iterate(campaignID string, params map[string]string) chan AppCampaignResultStruct {
	resp, err := ac.Get(campaignID, params)

	if err != nil {
		panic(err)
	}

	acrc := make(chan AppCampaignResultStruct)

	go appCampaignResults(acrc, resp, ac, campaignID)

	return acrc
}

// appCampaignResults feeds the results channel with items. While HasMore is true it makes new API requests
// and sends them through the channel. Once it is false it closes the channel.
func appCampaignResults(acrc chan AppCampaignResultStruct, resp *AppCampaignResultsResponse, acr *AppCampaignResults, campaignID string) {
	for {
		for _, item := range resp.Items {
			acrc <- item
		}
		if !resp.HasMore {
			close(acrc)
			return
		}
		params := map[string]string{
			"since": strconv.FormatInt(resp.LastTimestamp, 10),
		}

		resp, err := acr.Get(campaignID, params)

		if err != nil {
			panic(err)
		}

		appCampaignResults(acrc, resp, acr, campaignID)

		return
	}
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

// newAppCampaignResponse creates a response and unmarshals the JSON into a Struct.
func newAppCampaignResponse(data []byte) (*AppCampaignResponse, error) {
	response := &AppCampaignResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// newAppCampaignResultsResponse creates a response and unmarshals the JSON into a Struct.
func newAppCampaignResultsResponse(data []byte) (*AppCampaignResultsResponse, error) {
	response := &AppCampaignResultsResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
