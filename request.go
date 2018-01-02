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
	client *http.Client
}

// Get issues a GET request to the API and uses auth to set the authorization header.
func (r *request) get() ([]byte, error) {
	// Request also escapes whatever URL is passed here as string
	req, err := http.NewRequest(r.method, r.url(), nil)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	rfcdate := getRFC1123GMT(now)

	req.Header.Add("date", rfcdate)
	req.Header.Add("host", host)

	query := r.query()

	req.URL.RawQuery = query

	shortDate := getShortDate(now)
	shortDateTime := getShortDateTime(now)

	authHeader := r.auth.header(r.method, r.uri, query, rfcdate, host, shortDate, shortDateTime)

	req.Header.Add("authorization", authHeader)

	resp, err := r.client.Do(req)
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
// resource URI.
func (r *request) url() string {
	return fmt.Sprintf("%s://%s%s", scheme, host, r.uri)
}

// Query returns URL encoded query parameters using the params map that
// is passed in the Request.
func (r *request) query() string {
	v := url.Values{}
	for key, value := range r.params {
		v.Set(key, value)
	}
	return v.Encode()
}
