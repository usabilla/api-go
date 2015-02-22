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

func (au *auth) header(sts, shortDate string) string {
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
