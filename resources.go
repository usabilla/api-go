package gobilla

import (
	"fmt"
	"strconv"
)

// Canonical URI constants.
const (
	buttonURI   = "/live/website/button"
	campaignURI = "/live/website/campaign"
)

var (
	feedbackURI        = buttonURI + "/%s/feedback"
	campaignResultsURI = campaignURI + "/%s/results"
)

type resource struct {
	auth auth
}

// Buttons represents the button resource of Usabilla API.
type Buttons struct {
	resource
}

// Get function of Buttons resource returns all the buttons
// taking into account the specified query params.
// Accepted query params are:
// - limit string
func (b *Buttons) Get(params map[string]string) (*ButtonResponse, error) {
	request := Request{
		method: "GET",
		auth:   b.auth,
		uri:    buttonURI,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := ButtonResponse{}

	return response.unmarshal(data)
}

// Feedback encapsulates the feedback item resource.
func (b *Buttons) Feedback() *FeedbackItems {
	return &FeedbackItems{
		resource: resource{
			auth: b.auth,
		},
	}
}

// FeedbackItems represents the feedback item subresource of Usabilla API.
type FeedbackItems struct {
	resource
}

// Get function of FeedbackItem resource returns all the feedback items
// for a specific button, taking into account the passed query params.
// Accepted query params are:
// - since string (Time stamp)
func (f *FeedbackItems) Get(buttonID string, params map[string]string) (*FeedbackResponse, error) {
	uri := fmt.Sprintf(feedbackURI, buttonID)

	request := &Request{
		method: "GET",
		auth:   f.auth,
		uri:    uri,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := FeedbackResponse{}

	resp, err := response.unmarshal(data)

	return resp, err
}

// Iterate ...
func (f *FeedbackItems) Iterate(buttonID string, params map[string]string) chan FeedbackItem {
	resp, err := f.Get(buttonID, params)

	if err != nil {
		panic(err)
	}

	fic := make(chan FeedbackItem)

	go items(fic, resp, f, buttonID)

	return fic
}

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

		go items(fic, resp, f, buttonID)

		return
	}
}

// Campaigns represents the campaign resource of Usabilla API.
type Campaigns struct {
	resource
}

// Get function of Campaigns resource returns all the campaigns
// taking into account the passed query params.
// Accepted query params are:
// - limit string
// - since string (Time stamp)
func (c *Campaigns) Get(params map[string]string) (*CampaignResponse, error) {
	request := Request{
		method: "GET",
		auth:   c.auth,
		uri:    campaignURI,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := CampaignResponse{}

	return response.unmarshal(data)
}

// Results ...
func (c *Campaigns) Results() *CampaignResults {
	return &CampaignResults{
		resource: resource{
			auth: c.auth,
		},
	}
}

// CampaignResults represents the campaign result resource of Usabilla API.
type CampaignResults struct {
	resource
}

// Get function of CampaignResults resource returns all the campaign result items
// for a specific campaign, taking into account the passed query params.
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (r *CampaignResults) Get(campaignID string, params map[string]string) (*CampaignResultResponse, error) {
	uri := fmt.Sprintf(campaignResultsURI, campaignID)

	request := Request{
		method: "GET",
		auth:   r.auth,
		uri:    uri,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := CampaignResultResponse{}

	return response.unmarshal(data)
}

// Iterate ...
func (r *CampaignResults) Iterate(campaignID string, params map[string]string) chan CampaignResult {
	resp, err := r.Get(campaignID, params)

	if err != nil {
		panic(err)
	}

	crc := make(chan CampaignResult)

	go campaignResults(crc, resp, r, campaignID)

	return crc
}

func campaignResults(crc chan CampaignResult, resp *CampaignResultResponse, r *CampaignResults, campaignID string) {
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

		go campaignResults(crc, resp, r, campaignID)

		return
	}
}
