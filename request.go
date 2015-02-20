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
Request ...
*/
type Request struct {
	CanonicalURI string
}

/*
Get ...
*/
func (req *Request) Get(params map[string]string) ([]byte, error) {
	now := time.Now()
	rfcdate := now.Format(RFC1123GMT)
	shortDate := now.Format(ShortDate)
	shortDateTime := now.Format(ShortDateTime)

	method := "GET"

	url := fmt.Sprintf("%s://%s%s", scheme, host, req.CanonicalURI)

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("date", rfcdate)
	request.Header.Add("host", host)

	v := request.URL.Query()
	for key, value := range params {
		v.Set(key, value)
	}
	canonicalQueryString := v.Encode()
	request.URL.RawQuery = canonicalQueryString

	canonicalHeaders := fmt.Sprintf("date:%s\nhost:%s\n", rfcdate, host)

	signedHeaders := "date;host"

	payloadHash := hexHash([]byte(""))

	canonicalRequest := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		method,
		req.CanonicalURI,
		canonicalQueryString,
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
