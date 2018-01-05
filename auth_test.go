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
	"testing"

	"github.com/usabilla/api-go/internal"
)

var testData = map[string]string{
	"method":    "GET",
	"uri":       "/live/website/button",
	"query":     "limit=1",
	"rfcdate":   "Tue, 10 Feb 2015 23:00:00 GMT",
	"host":      "data.usabilla.com",
	"load":      "test",
	"shortDate": "20150223",
	"key":       "key",
	"secret":    "secret",
}

func fakeAuth() *auth {
	return &auth{
		key:    testData["key"],
		secret: testData["secret"],
	}
}

func Test_CredentialScope(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	spec.Expect(auth.credentialScope(testData["shortDate"])).ToEqual("20150223/usbl1_request")
}

func Test_SignedHeaders(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	spec.Expect(auth.signedHeaders()).ToEqual("date;host")
}

func Test_Payload(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	spec.Expect(auth.payload(testData["load"])).ToEqual(expected)
}

func Test_CanonicalHeaders(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	expected := "date:Tue, 10 Feb 2015 23:00:00 GMT\nhost:data.usabilla.com\n"
	spec.Expect(auth.canonicalHeaders(testData["rfcdate"], testData["host"])).ToEqual(expected)
}

func Test_CanonicalRequest(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	cr := auth.canonicalRequest(testData["method"], testData["uri"], testData["query"], testData["rfcdate"], testData["host"])
	expected := "GET\n/live/website/button\nlimit=1\ndate:Tue, 10 Feb 2015 23:00:00 GMT\nhost:data.usabilla.com\n\ndate;host\ne3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	spec.Expect(cr).ToEqual(expected)
}

func Test_HashedCanonincalRequest(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	cr := auth.hashedCanonicalRequest(testData["method"], testData["uri"], testData["query"], testData["rfcdate"], testData["host"])
	expected := "dd983da3893e6c9cccb4a2a11fe23a9380746cf0882ccd5f21304876d67cde8a"
	spec.Expect(cr).ToEqual(expected)
}

func Test_StringToSign(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	cr := auth.stringToSign(testData["method"], testData["uri"], testData["query"], testData["rfcdate"], testData["host"], testData["shortDate"], testData["shortDateTime"])
	expected := "USBL1-HMAC-SHA256\n\n20150223/usbl1_request\ndd983da3893e6c9cccb4a2a11fe23a9380746cf0882ccd5f21304876d67cde8a"
	spec.Expect(cr).ToEqual(expected)
}

func Test_Signature(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	signature := auth.signature(testData["method"], testData["uri"], testData["query"], testData["rfcdate"], testData["host"], testData["shortDate"], testData["shortDateTime"])
	expected := "a45df8012058d07b89e12cef34754029d2edcc80f930dca586e395b79ef2baf3"
	spec.Expect(signature).ToEqual(expected)
}

func Test_Header(t *testing.T) {
	spec := internal.Spec(t)
	auth := fakeAuth()
	header := auth.header(testData["method"], testData["uri"], testData["query"], testData["rfcdate"], testData["host"], testData["shortDate"], testData["shortDateTime"])
	expected := "USBL1-HMAC-SHA256 Credential=key/20150223/usbl1_request, SignedHeaders=date;host, Signature=a45df8012058d07b89e12cef34754029d2edcc80f930dca586e395b79ef2baf3"
	spec.Expect(header).ToEqual(expected)
}
