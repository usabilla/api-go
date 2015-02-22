package gobilla

import "fmt"

// Canonical URI constants.
const (
	buttonURI          = "/live/website/button"
	campaignURI        = "/live/website/campaign"
	feedbackURI        = "/feedback"
	campaignResultsURI = "/results"
)

type resource struct {
	auth auth
	uri  string
}

/*
Buttons represents the button resource of Usabilla API.
*/
type Buttons struct {
	resource
}

/*
Get function of Buttons resource returns all the buttons
taking into account the specified query params.

Accepted query params are:

- limit string
*/
func (b *Buttons) Get(params map[string]string) (*ButtonResponse, error) {
	request := Request{
		method: "GET",
		auth:   b.auth,
		uri:    b.uri,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := ButtonResponse{}

	return response.JSON(data)
}

// Feedback encapsulates the feedback item resource.
func (b *Buttons) Feedback() *FeedbackItems {
	uri := buttonURI + "/%s" + feedbackURI
	return &FeedbackItems{
		resource: resource{
			auth: b.auth,
			uri:  uri,
		},
	}
}

/*
FeedbackItems represents the feedback item resource of Usabilla API.
*/
type FeedbackItems struct {
	resource
}

/*
Get function of FeedbackItem resource returns all the feedback items
for a specific button, taking into account the passed query params.

Accepted query params are:

- since string (Time stamp)
*/
func (f *FeedbackItems) Get(buttonID string, params map[string]string) (*FeedbackResponse, error) {
	uri := fmt.Sprintf(f.uri, buttonID)

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

	return response.JSON(data)
}

// Campaigns represents the campaign resource of Usabilla API.
type Campaigns struct {
	resource
}

/*
Get function of Campaigns resource returns all the campaigns
taking into account the passed query params.

Accepted query params are:

- limit string
- since string (Time stamp)
*/
func (c *Campaigns) Get(params map[string]string) (*CampaignResponse, error) {
	request := Request{
		method: "GET",
		auth:   c.auth,
		uri:    c.uri,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := CampaignResponse{}

	return response.JSON(data)
}

// Results ...
func (c *Campaigns) Results() CampaignResults {
	uri := campaignURI + "/%s" + campaignResultsURI
	return CampaignResults{
		resource: resource{
			auth: c.auth,
			uri:  uri,
		},
	}
}

// CampaignResults represents the campaign result resource of Usabilla API.
type CampaignResults struct {
	resource
}

/*
Get function of CampaignResults resource returns all the campaign result items
for a specific campaign, taking into account the passed query params.

Accepted query params are:

- limit int
- since string (Time stamp)
*/
func (r *CampaignResults) Get(campaignID string, params map[string]string) (*CampaignResultResponse, error) {
	campaignURI := fmt.Sprintf(r.uri, campaignID)

	request := Request{
		method: "GET",
		auth:   r.auth,
		uri:    campaignURI,
		params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	response := CampaignResultResponse{}

	return response.JSON(data)
}
