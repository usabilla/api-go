package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/usabilla/gobilla/auth"
	"github.com/usabilla/gobilla/date"
)

const (
	scheme = "http"
	host   = "data.usabilla.com"
)

// Request is a request that the client makes to the API.
type Request struct {
	Auth   auth.Auth
	URI    string
	Method string
	Params map[string]string
	Client http.Client
}

// Get issues a GET request to the API and uses auth to set the authorization header.
func (r *Request) Get() ([]byte, error) {
	// Request also escapes whatever URL is passed here as string
	request, err := http.NewRequest(r.Method, r.URL(), nil)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	rfcdate := date.GetRFC1123GMT(now)

	request.Header.Add("date", rfcdate)
	request.Header.Add("host", host)

	query := r.Query()

	request.URL.RawQuery = query

	shortDate := date.GetShortDate(now)
	shortDateTime := date.GetShortDateTime(now)

	authHeader := r.Auth.Header(r.Method, r.URI, query, rfcdate, host, shortDate, shortDateTime)

	request.Header.Add("authorization", authHeader)

	resp, err := r.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// URL returns the request URL using the scheme, host, and the additional
// resource URI
func (r *Request) URL() string {
	return fmt.Sprintf("%s://%s%s", scheme, host, r.URI)
}

// Query returns URL encoded query parameters using the params map that
// is passed in the Request
func (r *Request) Query() string {
	v := url.Values{}
	for key, value := range r.Params {
		v.Set(key, value)
	}
	return v.Encode()
}
