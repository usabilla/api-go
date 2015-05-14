package auth

import "fmt"

const (
	algorithm  = "USBL1-HMAC-SHA256"
	terminator = "usbl1_request"
	startor    = "USBL1"
)

// Auth holds the key and secret information and provides methods that
// encapsulate the whole request signing process of the API
type Auth struct {
	Key, Secret string
}

// Header creates the authorization header.
func (au *Auth) Header(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	return fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		au.Key,
		au.credentialScope(shortDate),
		au.signedHeaders(),
		au.signature(method, uri, query, rfcdate, host, shortDate, shortDateTime),
	)
}

// Create the credential scope that includes the short date format and termination string.
func (au *Auth) credentialScope(shortDate string) string {
	return shortDate + "/" + terminator
}

// Create a signature using the string to sign.
func (au *Auth) signature(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	dig := keyedHash([]byte(startor+au.Secret), []byte(shortDate))

	dig = keyedHash(dig, []byte(terminator))

	sts := au.stringToSign(method, uri, query, rfcdate, host, shortDate, shortDateTime)

	return hexKeyedHash(dig, []byte(sts))
}

// Return the signed headers.
func (au *Auth) signedHeaders() string {
	return "date;host"
}

// Create a hexademical hash of the payload.
func (au *Auth) payload(load string) string {
	return hexHash([]byte(load))
}

// Create a hexademical hash of the canonical request.
func (au *Auth) hashedCanonicalRequest(method, uri, query, rfcdate, host string) string {
	return hexHash([]byte(au.canonicalRequest(method, uri, query, rfcdate, host)))
}

// Create the string to be used for signing.
func (au *Auth) stringToSign(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		algorithm,
		shortDateTime,
		au.credentialScope(shortDate),
		au.hashedCanonicalRequest(method, uri, query, rfcdate, host),
	)
}

// Create a canonical format of the request.
func (au *Auth) canonicalRequest(method, uri, query, rfcdate, host string) string {
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
func (au *Auth) canonicalHeaders(rfcdate, host string) string {
	return fmt.Sprintf("date:%s\nhost:%s\n", rfcdate, host)
}
