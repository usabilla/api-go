package l4a

import (
	"fmt"
	"strconv"

	"github.com/usabilla/gobilla/request"
	"github.com/usabilla/gobilla/resource"
)

// Canonical URI constants.
const (
	appsURI = "/live/apps"
)

var (
	feedbackURI = appsURI + "/%s/feedback"
)

// Apps represents the app resource of Usabilla API.
type Apps struct {
	resource.Resource
}

// Get function of Apps resource returns all apps
// taking into account the specified query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (a *Apps) Get(params map[string]string) (*AppResponse, error) {
	request := request.Request{
		Method: "GET",
		Auth:   a.Auth,
		URI:    appsURI,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewAppResponse(data)
}

// Feedback encapsulates the app feedback item resource.
func (a *Apps) Feedback() *AppFeedbackItems {
	return &AppFeedbackItems{
		Resource: resource.Resource{
			Auth: a.Auth,
		},
	}
}

// AppFeedbackItems represents the apps feedback item subresource of Usabilla API.
type AppFeedbackItems struct {
	resource.Resource
}

// Get function of AppFeedbackItem resource returns all the feedback items
// for a specific app, taking into account the provided query params.
//
// Accepted query params are:
// - limit int
// - since string (Time stamp)
func (af *AppFeedbackItems) Get(appID string, params map[string]string) (*AppFeedbackResponse, error) {
	uri := fmt.Sprintf(feedbackURI, appsURI, appID)

	request := &request.Request{
		Method: "GET",
		Auth:   af.Auth,
		URI:    uri,
		Params: params,
	}

	data, err := request.Get()
	if err != nil {
		panic(err)
	}

	return NewAppFeedbackResponse(data)
}

// Iterate uses an AppFeedbackItem channel which transparently uses the HasMore field to fire
// a new api request once all items have been consumed on the channel
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
// When HasMore is false, we close the channel
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

		go appItems(afic, resp, af, appID)

		return
	}
}
