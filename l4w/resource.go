package l4w

import (
	"fmt"
	"strconv"

	"github.com/usabilla/gobilla/request"
	"github.com/usabilla/gobilla/resource"
)

// Canonical URI constants.
const (
	websitesURI = "/live/websites"
	buttonURI   = websitesURI + "/button"
	campaignURI = websitesURI + "/campaign"
)

var (
	feedbackURI        = buttonURI + "/%s/feedback"
	campaignResultsURI = campaignURI + "/%s/results"
	campaignStatsURI   = campaignURI + "/%s/stats"
)

// Buttons represents the button resource of Usabilla API.
type Buttons struct {
	resource.Resource
}

// Get function of Buttons resource returns all the buttons
// taking into account the specified query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (b *Buttons) Get(params map[string]string) (*ButtonResponse, error) {
	request := request.Request{
		Method: "GET",
		Auth:   b.Auth,
		URI:    buttonURI,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewButtonResponse(data)
}

// Feedback encapsulates the feedback item resource.
func (b *Buttons) Feedback() *FeedbackItems {
	return &FeedbackItems{
		Resource: resource.Resource{
			Auth: b.Auth,
		},
	}
}

// FeedbackItems represents the feedback item subresource of Usabilla API.
type FeedbackItems struct {
	resource.Resource
}

// Get function of FeedbackItem resource returns all the feedback items
// for a specific button, taking into account the provided query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (f *FeedbackItems) Get(buttonID string, params map[string]string) (*FeedbackResponse, error) {
	uri := fmt.Sprintf(feedbackURI, buttonID)

	request := &request.Request{
		Method: "GET",
		Auth:   f.Auth,
		URI:    uri,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewFeedbackResponse(data)
}

// Iterate uses a FeedbackItem channel which transparently uses the HasMore field to fire
// a new api request once all items have been consumed on the channel
func (f *FeedbackItems) Iterate(buttonID string, params map[string]string) chan FeedbackItem {
	resp, err := f.Get(buttonID, params)

	if err != nil {
		panic(err)
	}

	fic := make(chan FeedbackItem)

	go items(fic, resp, f, buttonID)

	return fic
}

// items feeds a feedback item channel with items
//
// while hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items
//
// when HasMore is false, we close the channel
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
	resource.Resource
}

// Get function of Campaigns resource returns all the campaigns
// taking into account the provided query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (c *Campaigns) Get(params map[string]string) (*CampaignResponse, error) {
	request := request.Request{
		Method: "GET",
		Auth:   c.Auth,
		URI:    campaignURI,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewCampaignResponse(data)
}

// Results encapsulates the campaign results resource.
func (c *Campaigns) Results() *CampaignResults {
	return &CampaignResults{
		Resource: resource.Resource{
			Auth: c.Auth,
		},
	}
}

// CampaignResults represents the campaign result resource of Usabilla API.
type CampaignResults struct {
	resource.Resource
}

// Get function of CampaignResults resource returns all the campaign result items
// for a specific campaign, taking into account the provided query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (r *CampaignResults) Get(campaignID string, params map[string]string) (*CampaignResultResponse, error) {
	uri := fmt.Sprintf(campaignResultsURI, campaignID)

	request := request.Request{
		Method: "GET",
		Auth:   r.Auth,
		URI:    uri,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewCampaignResultResponse(data)
}

// Iterate uses a CampaignResult channel which transparently uses the HasMore field to fire
// a new api request once all results have been consumed on the channel
func (r *CampaignResults) Iterate(campaignID string, params map[string]string) chan CampaignResult {
	resp, err := r.Get(campaignID, params)

	if err != nil {
		panic(err)
	}

	crc := make(chan CampaignResult)

	go campaignResults(crc, resp, r, campaignID)

	return crc
}

// campaignResults feeds a campaign results channel with items
//
// while hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items
//
// when HasMore is false, we close the channel
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

// Stats encapsulates the campaign statistics resource.
func (c *Campaigns) Stats() *CampaignStats {
	return &CampaignStats{
		Resource: resource.Resource{
			Auth: c.Auth,
		},
	}
}

// CampaignStats represents the campaign statistics resource of Usabilla API.
type CampaignStats struct {
	resource.Resource
}

// Get function of CampaignStats resource returns the campaign statistics
// for a specific campaign, taking into account the provided query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (cs *CampaignStats) Get(campaignID string, params map[string]string) (*CampaignStatsResponse, error) {
	uri := fmt.Sprintf(campaignStatsURI, campaignID)

	request := request.Request{
		Method: "GET",
		Auth:   cs.Auth,
		URI:    uri,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewCampaignStatsResponse(data)
}

// Iterate uses a CampaignStat channel which transparently uses the HasMore field to fire
// a new api request once all stats items have been consumed on the channel
func (cs *CampaignStats) Iterate(campaignID string, params map[string]string) chan CampaignStat {
	resp, err := cs.Get(campaignID, params)

	if err != nil {
		panic(err)
	}

	csc := make(chan CampaignStat)

	go campaignStats(csc, resp, cs, campaignID)

	return csc
}

// campagnStats feeds a campaign statistics channel with items
//
// while hasMore is true and all items have been consumed in the channel
// a new request is fired using the since parameter of the response, to
// retrieve new items
//
// when HasMore is false, we close the channel
func campaignStats(csc chan CampaignStat, resp *CampaignStatsResponse, cs *CampaignStats, campaignID string) {
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

		go campaignStats(csc, resp, cs, campaignID)

		return
	}
}
