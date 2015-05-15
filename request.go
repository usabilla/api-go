package gobilla

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	scheme = "http"
	host   = "data.usabilla.com"
)

// Request is a request that the client makes to the API.
type request struct {
	auth   auth
	uri    string
	method string
	params map[string]string
	client http.Client
}

// Get issues a GET request to the API and uses auth to set the authorization header.
func (r *request) get() ([]byte, error) {
	// Request also escapes whatever URL is passed here as string
	request, err := http.NewRequest(r.method, r.url(), nil)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	rfcdate := getRFC1123GMT(now)

	request.Header.Add("date", rfcdate)
	request.Header.Add("host", host)

	query := r.query()

	request.URL.RawQuery = query

	shortDate := getShortDate(now)
	shortDateTime := getShortDateTime(now)

	authHeader := r.auth.header(r.method, r.uri, query, rfcdate, host, shortDate, shortDateTime)

	request.Header.Add("authorization", authHeader)

	resp, err := r.client.Do(request)
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
func (r *request) url() string {
	return fmt.Sprintf("%s://%s%s", scheme, host, r.uri)
}

// Query returns URL encoded query parameters using the params map that
// is passed in the Request
func (r *request) query() string {
	v := url.Values{}
	for key, value := range r.params {
		v.Set(key, value)
	}
	return v.Encode()
}
