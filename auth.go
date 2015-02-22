package gobilla

import "fmt"

const (
	algorithm  = "USBL1-HMAC-SHA256"
	terminator = "usbl1_request"
	startor    = "USBL1"
)

type auth struct {
	key, secret string
}

// Create the authorization header.
func (au *auth) header(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	sts := au.stringToSign(method, uri, query, rfcdate, host, shortDate, shortDateTime)
	return fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		au.key,
		au.credentialScope(shortDate),
		au.signedHeaders(),
		au.signature(sts, shortDate),
	)
}

// Create the credential scope that includes the short date format and termination string.
func (au *auth) credentialScope(shortDate string) string {
	return shortDate + "/" + terminator
}

// Create a signature using the string to sign.
func (au *auth) signature(sts, shortDate string) string {
	dig := keyedHash([]byte(startor+au.secret), []byte(shortDate))

	dig = keyedHash(dig, []byte(terminator))

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
