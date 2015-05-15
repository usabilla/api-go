package gobilla

import (
	"testing"

	"github.com/usabilla/gobilla/internal"
)

func mockRequest(method, uri string, params map[string]string) *request {
	return &request{
		auth:   auth{},
		uri:    uri,
		method: method,
		params: params,
	}
}

func Test_URL(t *testing.T) {
	spec := internal.Spec(t)
	r := mockRequest("GET", "/live/websites/button", nil)
	expected := "http://data.usabilla.com/live/websites/button"
	spec.Expect(r.url()).ToEqual(expected)
}

func Test_Query(t *testing.T) {
	spec := internal.Spec(t)
	r := mockRequest("GET", "/", map[string]string{"name": "value"})
	expected := "name=value"
	spec.Expect(r.query()).ToEqual(expected)
}
