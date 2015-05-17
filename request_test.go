/*
Copyright (c) 2015 Usabilla

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

package gobilla

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/usabilla/gobilla/internal"
)

func mockRequest(method, uri string, params map[string]string, client *http.Client) *request {
	return &request{
		auth:   auth{},
		uri:    uri,
		method: method,
		params: params,
		client: client,
	}
}

func mockServerClient(code int, method, body string, params map[string]string) (*httptest.Server, *request) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	request := mockRequest(method, server.URL, params, httpClient)

	return server, request
}

func Test_URL(t *testing.T) {
	spec := internal.Spec(t)
	r := mockRequest("GET", "/live/websites/button", nil, nil)
	expected := "http://data.usabilla.com/live/websites/button"
	spec.Expect(r.url()).ToEqual(expected)
}

func Test_Query(t *testing.T) {
	spec := internal.Spec(t)
	r := mockRequest("GET", "/", map[string]string{"name": "value"}, nil)
	expected := "name=value"
	spec.Expect(r.query()).ToEqual(expected)
}

func Test_Get_Successful(t *testing.T) {
	exampleResponse := `{"count": 100, "hasMore": true, "lastTimestamp": 1431867114}`
	server, request := mockServerClient(200, "GET", exampleResponse, nil)
	defer server.Close()

	mockResponse := &response{}

	data, err := request.get()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	err = json.Unmarshal(data, &mockResponse)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	expectedResponse := &response{
		Count:         100,
		HasMore:       true,
		LastTimestamp: 1431867114,
	}

	spec := internal.Spec(t)
	spec.Expect(reflect.DeepEqual(mockResponse, expectedResponse)).ToEqual(true)
}
