package request

import (
	"testing"

	"github.com/usabilla/gobilla/auth"
	"github.com/usabilla/gobilla/internal"
)

func mockRequest(method, uri string, params map[string]string) *Request {
	return &Request{
		Auth:   auth.Auth{},
		URI:    uri,
		Method: method,
		Params: params,
	}
}

func Test_URL(t *testing.T) {
	spec := internal.Spec(t)
	r := mockRequest("GET", "/live/websites/button", nil)
	expected := "http://data.usabilla.com/live/websites/button"
	spec.Expect(r.URL()).ToEqual(expected)
}

func Test_Query(t *testing.T) {
	spec := internal.Spec(t)
	r := mockRequest("GET", "/", map[string]string{"name": "value"})
	expected := "name=value"
	spec.Expect(r.Query()).ToEqual(expected)
}
