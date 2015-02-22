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

func (au *auth) credentialScope(shortDate string) string {
	return shortDate + "/" + terminator
}

func (au *auth) signature(sts, shortDate string) string {
	dig := keyedHash([]byte(startor+au.secret), []byte(shortDate))

	dig = keyedHash(dig, []byte(terminator))

	return hexKeyedHash(dig, []byte(sts))
}

func (au *auth) signedHeaders() string {
	return "date;host"
}

func (au *auth) payload(load string) string {
	return hexHash([]byte(load))
}

func (au *auth) hashedCanonicalRequest(method, uri, query, rfcdate, host string) string {
	return hexHash([]byte(au.canonicalRequest(method, uri, query, rfcdate, host)))
}

func (au *auth) stringToSign(method, uri, query, rfcdate, host, shortDate, shortDateTime string) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		algorithm,
		shortDateTime,
		au.credentialScope(shortDate),
		au.hashedCanonicalRequest(method, uri, query, rfcdate, host),
	)
}

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

func (au *auth) canonicalHeaders(rfcdate, host string) string {
	return fmt.Sprintf("date:%s\nhost:%s\n", rfcdate, host)
}
