package gobilla

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	scheme     = "http"
	host       = "data.usabilla.com"
	algorithm  = "USBL1-HMAC-SHA256"
	terminator = "usbl1_request"
)

/*
APIRequest ...
*/
type APIRequest struct {
	Request      *http.Request
	CanonicalURI string
}

/*
Get ...
*/
func (req *APIRequest) Get(params map[string]string) ([]byte, error) {
	now := time.Now()
	rfcdate := now.Format(RFC1123GMT)
	shortDate := now.Format(ShortDate)
	shortDateTime := now.Format(ShortDateTime)

	method := "GET"

	request, err := http.NewRequest(method, req.url(), nil)
	if err != nil {
		return nil, err
	}
	req.Request = request

	req.Request.Header.Add("date", rfcdate)
	req.Request.Header.Add("host", host)

	canonicalQuery := req.canonicalQueryString(params)

	canonicalHeaders := fmt.Sprintf("date:%s\nhost:%s\n", rfcdate, host)

	signedHeaders := "date;host"

	payloadHash := hexHash([]byte(""))

	canonicalRequest := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		method,
		req.CanonicalURI,
		canonicalQuery,
		canonicalHeaders,
		signedHeaders,
		payloadHash,
	)

	credentialScope := shortDate + "/" + terminator

	hashedCanonicalRequest := hexHash([]byte(canonicalRequest))

	stringToSign := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		algorithm,
		shortDateTime,
		credentialScope,
		hashedCanonicalRequest,
	)

	digest2 := keyedHash([]byte("USBL1"+secret), []byte(shortDate))

	digest3 := keyedHash(digest2, []byte(terminator))

	signature := hexKeyedHash(digest3, []byte(stringToSign))

	authorizationHeader := fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		key,
		credentialScope,
		signedHeaders,
		signature,
	)

	request.Header.Add("authorization", authorizationHeader)

	client := http.Client{}
	resp, err := client.Do(request)
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

func (req *APIRequest) url() string {
	return fmt.Sprintf("%s://%s%s", scheme, host, req.CanonicalURI)
}

func (req *APIRequest) canonicalQueryString(params map[string]string) string {
	v := req.Request.URL.Query()
	for key, value := range params {
		v.Set(key, value)
	}
	canonicalQuery := v.Encode()
	req.Request.URL.RawQuery = canonicalQuery
	return canonicalQuery
}
