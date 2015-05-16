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

import "fmt"

const (
	algorithm  = "USBL1-HMAC-SHA256"
	terminator = "usbl1_request"
	startor    = "USBL1"
)

// Auth holds the key and secret information and provides methods that
// encapsulate the whole request signing process of the API.
type auth struct {
	key, secret string
}

// Header creates the authorization header.
func (au *auth) header(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	return fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		au.key,
		au.credentialScope(shortDate),
		au.signedHeaders(),
		au.signature(method, uri, query, rfcdate, host, shortDate, shortDateTime),
	)
}

// Create the credential scope that includes the short date format and termination string.
func (au *auth) credentialScope(shortDate string) string {
	return shortDate + "/" + terminator
}

// Create a signature using the string to sign.
func (au *auth) signature(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	dig := keyedHash([]byte(startor+au.secret), []byte(shortDate))

	dig = keyedHash(dig, []byte(terminator))

	sts := au.stringToSign(method, uri, query, rfcdate, host, shortDate, shortDateTime)

	return hexKeyedHash(dig, []byte(sts))
}

// Return the signed headers.
func (au *auth) signedHeaders() string {
	return "date;host"
}

// Create a hexademical hash of the payload.
func (au *auth) payload(load string) string {
	return hexHash([]byte(load))
}

// Create a hexademical hash of the canonical request.
func (au *auth) hashedCanonicalRequest(method, uri, query, rfcdate, host string) string {
	return hexHash([]byte(au.canonicalRequest(method, uri, query, rfcdate, host)))
}

// Create the string to be used for signing.
func (au *auth) stringToSign(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		algorithm,
		shortDateTime,
		au.credentialScope(shortDate),
		au.hashedCanonicalRequest(method, uri, query, rfcdate, host),
	)
}

// Create a canonical format of the request.
func (au *auth) canonicalRequest(method, uri, query, rfcdate, host string) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		method,
		uri,
		query,
		au.canonicalHeaders(rfcdate, host),
		au.signedHeaders(),
		au.payload(""),
	)
}

// Create a canonical format of the headers.
func (au *auth) canonicalHeaders(rfcdate, host string) string {
	return fmt.Sprintf("date:%s\nhost:%s\n", rfcdate, host)
}
