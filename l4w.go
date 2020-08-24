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
	"time"
)

// Canonical URI constants.
const (
	websitesURI = "/live/websites"
	buttonURI   = websitesURI + "/button"
	campaignURI = websitesURI + "/campaign"
	inpageURI   = websitesURI + "/inpage"
)

var (
	feedbackURI        = buttonURI + "/%s/feedback"
	campaignResultsURI = campaignURI + "/%s/results"
	campaignStatsURI   = campaignURI + "/%s/stats"
	inpageFeedbackURI  = inpageURI + "/%s/feedback"
)

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
type CampaignStruct struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	ButtonID    string    `json:"buttonId"`
	AnalyticsID string    `json:"analyticsId"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
}

// CampaignResult represents a campaign result item.
type CampaignResultStruct struct {
	ID         string                 `json:"id"`
	UserAgent  string                 `json:"userAgent"`
	Location   string                 `json:"location"`
	Date       time.Time              `json:"date"`
	CampaignID string                 `json:"campaignId"`
	CustomData map[string]string      `json:"customData"`
	Data       map[string]interface{} `json:"data"`
	URL        string                 `json:"url"`
	Time       float64                `json:"time"`
}

// CampaignStat represents a campaign statistics item.
type CampaignStatStruct struct {
	ID         string `json:"id"`
	Completed  int    `json:"completed"`
	Conversion int    `json:"conversion"`
	Views      int    `json:"views"`
}

type InpageWidgetStruct struct {
	ID   string    `json:"id"`
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}

type InpageWidgetFeedbackStruct struct {
	ID         string                 `json:"id"`
	Date       time.Time              `json:"date"`
	Data       map[string]interface{} `json:"data"`
	CustomData map[string]interface{} `json:"customData"`
	Widget_ID  string                 `json:"widgetId"`
	Rating     float64                `json:"rating"`
	Mood       int                    `json:"mood"`
	Nps        int                    `json:"nps,omitempty"`
	Comment    string                 `json:"comment"`
	UserAgent  string                 `json:"userAgent"`
	Geo        GeoLoc                 `json:"geo"`
	Url        string                 `json:"url"`
}

type GeoLoc struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}

// ButtonResponse is a response that contains button data.
type ButtonResponse struct {
	response
	Items []Button `json:"items"`
}

// FeedbackResponse is a response that contains feedback item data.
type FeedbackResponse struct {
	response
	Items []FeedbackItem `json:"items"`
}

// CampaignResponse is a response that contains campaign data.
type CampaignResponse struct {
	response
	Items []CampaignStruct `json:"items"`
}

// CampaignResultResponse is a response that contains campaign result data.
type CampaignResultResponse struct {
	response
	Items []CampaignResultStruct `json:"items"`
}

// CampaignStatsResponse is a response that contains campaign statistics data.
type CampaignStatsResponse struct {
	response
	Items []CampaignStatStruct `json:"items"`
}

// InpageWidgetResponse is a response that lists inpage widgets.
type InpageWidgetResponse struct {
	response
	Items []InpageWidgetStruct `json:"items"`
}

// InpageWidgetFeedbackResponse is the response with in-page feedback data.
type InpageWidgetFeedbackResponse struct {
	response
	Items []InpageWidgetFeedbackStruct `json:"items"`
}

// Buttons represents the button resource of Usabilla API.
type Buttons struct {
	resource
	client *http.Client
}

// FeedbackItems represents the feedback item subresource of Usabilla API.
type FeedbackItems struct {
	resource
	client *http.Client
}

// Campaigns represents the campaign resource of Usabilla API.
type Campaigns struct {
	resource
	client *http.Client
}

// CampaignResults represents the campaign result resource of Usabilla API.
type CampaignResults struct {
	resource
	client *http.Client
}

// CampaignStats represents the campaign statistics resource of Usabilla API.
type CampaignStats struct {
	resource
	client *http.Client
}

// InpageWidgets represents the inpage widgets resource of Usabilla API.
type InpageWidgets struct {
	resource
	client *http.Client
}

// InpageWidgetFeedbackItems represents the inpage feedbacks resource of Usabilla API.
type InpageWidgetFeedbackItems struct {
	resource
	client *http.Client
}

// NewButtonResponse creates a button response and unmarshals json API
// button response to Go struct.
func NewButtonResponse(data []byte) (*ButtonResponse, error) {
	response := &ButtonResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewFeedbackResponse creates a feedback response and unmarshals json API
// feedback items response to Go struct.
func NewFeedbackResponse(data []byte) (*FeedbackResponse, error) {
	response := &FeedbackResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewCampaignResponse creates a campaign response and unmarshals json API
// campaign response to Go struct.
func NewCampaignResponse(data []byte) (*CampaignResponse, error) {
	response := &CampaignResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewCampaignResultResponse creates a new campaign result response and unmarshals json API
// campaign results response to Go struct.
func NewCampaignResultResponse(data []byte) (*CampaignResultResponse, error) {
	response := &CampaignResultResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewCampaignStatsResponse creates a new campaign statistics response and unmarshals json API
// campaign statistics response to Go struct.
func NewCampaignStatsResponse(data []byte) (*CampaignStatsResponse, error) {
	response := &CampaignStatsResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewInpageWidgetResponse creates a inpage response and unmarshals json API
// inpage response to Go struct.
func NewInpageWidgetResponse(data []byte) (*InpageWidgetResponse, error) {
	response := &InpageWidgetResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewInpageWidgetFeedbackResponse creates a new inpage feedback response and unmarshals json API
// inpage feedback response to Go struct.
func NewInpageWidgetFeedbackResponse(data []byte) (*InpageWidgetFeedbackResponse, error) {
	response := &InpageWidgetFeedbackResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Get function of Buttons resource returns all the buttons
// taking into account the specified query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (b *Buttons) Get(params map[string]string) (*ButtonResponse, error) {
	request := request{
		method: "GET",
		auth:   b.auth,
		uri:    buttonURI,
		params: params,
		client: b.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewButtonResponse(data)
}

// Feedback encapsulates the feedback item resource.
func (b *Buttons) Feedback() *FeedbackItems {
	return &FeedbackItems{
		resource: resource{
			auth: b.auth,
		},
		client: b.client,
	}
}

// Get function of FeedbackItem resource returns all the feedback items
// for a specific button, taking into account the provided query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (f *FeedbackItems) Get(buttonID string, params map[string]string) (*FeedbackResponse, error) {
	uri := fmt.Sprintf(feedbackURI, buttonID)

	request := &request{
		method: "GET",
		auth:   f.auth,
		uri:    uri,
		params: params,
		client: f.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewFeedbackResponse(data)
}

// Iterate uses a FeedbackItem channel which transparently uses the HasMore field to fire
// a new api request once all items have been consumed on the channel.
func (f *FeedbackItems) Iterate(buttonID string, params map[string]string) chan FeedbackItem {
	resp, err := f.Get(buttonID, params)

	if err != nil {
		panic(err)
	}

	fic := make(chan FeedbackItem)

	go items(fic, resp, f, buttonID)

	return fic
}

// items feeds a feedback item channel with items.
//
// While hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items.
//
// When HasMore is false, we close the channel.
func items(fic chan FeedbackItem, resp *FeedbackResponse, f *FeedbackItems, buttonID string) {
	for {
		for _, item := range resp.Items {
			fic <- item
		}
		if !resp.HasMore {
			close(fic)
			return
		}
		params := map[string]string{
			"since": strconv.FormatInt(resp.LastTimestamp, 10),
		}

		resp, err := f.Get(buttonID, params)

		if err != nil {
			panic(err)
		}

		items(fic, resp, f, buttonID)

		return
	}
}

// Get function of Campaigns resource returns all the campaigns
// taking into account the provided query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (c *Campaigns) Get(params map[string]string) (*CampaignResponse, error) {
	request := request{
		method: "GET",
		auth:   c.auth,
		uri:    campaignURI,
		params: params,
		client: c.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewCampaignResponse(data)
}

// Results encapsulates the campaign results resource.
func (c *Campaigns) Results() *CampaignResults {
	return &CampaignResults{
		resource: resource{
			auth: c.auth,
		},
		client: c.client,
	}
}

// Get function of CampaignResults resource returns all the campaign result items
// for a specific campaign, taking into account the provided query parameters.
//
// Valid query params are:
//  limit int
//  since string (Time stamp)
func (r *CampaignResults) Get(campaignID string, params map[string]string) (*CampaignResultResponse, error) {
	uri := fmt.Sprintf(campaignResultsURI, campaignID)

	request := request{
		method: "GET",
		auth:   r.auth,
		uri:    uri,
		params: params,
		client: r.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewCampaignResultResponse(data)
}

// Iterate uses a CampaignResult channel which transparently uses the HasMore field to fire
// a new api request once all results have been consumed on the channel
func (r *CampaignResults) Iterate(campaignID string, params map[string]string) chan CampaignResultStruct {
	resp, err := r.Get(campaignID, params)

	if err != nil {
		panic(err)
	}

	crc := make(chan CampaignResultStruct)

	go yieldCampaignResults(crc, resp, r, campaignID)

	return crc
}

// campaignResults feeds a campaign results channel with items
//
// while hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items
//
// when HasMore is false, we close the channel
func yieldCampaignResults(crc chan CampaignResultStruct, resp *CampaignResultResponse, r *CampaignResults, campaignID string) {
	for {
		for _, item := range resp.Items {
			crc <- item
		}
		if !resp.HasMore {
			close(crc)
			return
		}
		params := map[string]string{
			"since": strconv.FormatInt(resp.LastTimestamp, 10),
		}

		resp, err := r.Get(campaignID, params)

		if err != nil {
			panic(err)
		}

		yieldCampaignResults(crc, resp, r, campaignID)

		return
	}
}

// Stats encapsulates the campaign statistics resource.
func (c *Campaigns) Stats() *CampaignStats {
	return &CampaignStats{
		resource: resource{
			auth: c.auth,
		},
		client: c.client,
	}
}

// Get function of CampaignStats resource returns the campaign statistics
// for a specific campaign, taking into account the provided query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (cs *CampaignStats) Get(campaignID string, params map[string]string) (*CampaignStatsResponse, error) {
	uri := fmt.Sprintf(campaignStatsURI, campaignID)

	request := request{
		method: "GET",
		auth:   cs.auth,
		uri:    uri,
		params: params,
		client: cs.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewCampaignStatsResponse(data)
}

// Iterate uses a CampaignStat channel which transparently uses the HasMore field to fire
// a new api request once all stats items have been consumed on the channel.
func (cs *CampaignStats) Iterate(campaignID string, params map[string]string) chan CampaignStatStruct {
	resp, err := cs.Get(campaignID, params)

	if err != nil {
		panic(err)
	}

	csc := make(chan CampaignStatStruct)

	go yieldCampaignStats(csc, resp, cs, campaignID)

	return csc
}

// campagnStats feeds a campaign statistics channel with items
//
// while hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items
//
// when HasMore is false, we close the channel
func yieldCampaignStats(csc chan CampaignStatStruct, resp *CampaignStatsResponse, cs *CampaignStats, campaignID string) {
	for {
		for _, item := range resp.Items {
			csc <- item
		}
		if !resp.HasMore {
			close(csc)
			return
		}
		params := map[string]string{
			"since": strconv.FormatInt(resp.LastTimestamp, 10),
		}

		resp, err := cs.Get(campaignID, params)

		if err != nil {
			panic(err)
		}

		yieldCampaignStats(csc, resp, cs, campaignID)

		return
	}
}

////////////// Inpage //////////////
// Get function of Inpage resource returns all the Inpage Widgets
// taking into account the provided query parameters.
//
// Valid query parameters are:
//  limit int
//  since string (Time stamp)
func (c *InpageWidgets) Get(params map[string]string) (*InpageWidgetResponse, error) {
	request := request{
		method: "GET",
		auth:   c.auth,
		uri:    inpageURI,
		params: params,
		client: c.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewInpageWidgetResponse(data)
}

// Results encapsulates the inpage feedback resource.
func (c *InpageWidgets) Feedback() *InpageWidgetFeedbackItems {
	return &InpageWidgetFeedbackItems{
		resource: resource{
			auth: c.auth,
		},
		client: c.client,
	}
}

// Get function of CampaignResults resource returns all the campaign result items
// for a specific campaign, taking into account the provided query parameters.
//
// Valid query params are:
//  limit int
//  since string (Time stamp)
func (r *InpageWidgetFeedbackItems) Get(inpageID string, params map[string]string) (*InpageWidgetFeedbackResponse, error) {
	uri := fmt.Sprintf(inpageFeedbackURI, inpageID)

	request := request{
		method: "GET",
		auth:   r.auth,
		uri:    uri,
		params: params,
		client: r.client,
	}

	data, err := request.get()
	if err != nil {
		panic(err)
	}

	return NewInpageWidgetFeedbackResponse(data)
}

// Iterate uses a InpageWidgetFeedbackItems channel which transparently uses the HasMore field to fire
// a new api request once all results have been consumed on the channel
func (r *InpageWidgetFeedbackItems) Iterate(inpageID string, params map[string]string) chan InpageWidgetFeedbackStruct {
	resp, err := r.Get(inpageID, params)

	if err != nil {
		panic(err)
	}

	crc := make(chan InpageWidgetFeedbackStruct)

	go yieldInpageWidgetFeedbackItems(crc, resp, r, inpageID)

	return crc
}

// yieldInpageWidgetFeedbackItems feeds a inpage results channel with items
//
// while hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items
//
// when HasMore is false, we close the channel
func yieldInpageWidgetFeedbackItems(crc chan InpageWidgetFeedbackStruct, resp *InpageWidgetFeedbackResponse, r *InpageWidgetFeedbackItems, inpageID string) {
	for {
		for _, item := range resp.Items {
			crc <- item
		}
		if !resp.HasMore {
			close(crc)
			return
		}
		params := map[string]string{
			"since": strconv.FormatInt(resp.LastTimestamp, 10),
		}

		resp, err := r.Get(inpageID, params)

		if err != nil {
			panic(err)
		}

		yieldInpageWidgetFeedbackItems(crc, resp, r, inpageID)
		return
	}
}
