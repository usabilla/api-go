package gobilla

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	scheme = "http"
	host   = "data.usabilla.com"
)

/*
Request ...
*/
type Request struct {
	auth   auth
	uri    string
	method string
	params map[string]string
}

/*
Get ...
*/
func (r *Request) Get() ([]byte, error) {
	now := time.Now()
	rfcdate := now.Format(RFC1123GMT)
	shortDate := now.Format(ShortDate)
	shortDateTime := now.Format(ShortDateTime)

	request, err := http.NewRequest(r.method, r.url(), nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("date", rfcdate)
	request.Header.Add("host", host)

	query := r.query()

	request.URL.RawQuery = query

	authHeader := r.auth.header(r.method, r.uri, query, rfcdate, host, shortDate, shortDateTime)

	request.Header.Add("authorization", authHeader)

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

func (r *Request) url() string {
	return fmt.Sprintf("%s://%s%s", scheme, host, r.uri)
}

func (r *Request) query() string {
	v := url.Values{}
	for key, value := range r.params {
		v.Set(key, value)
	}
	return v.Encode()
}
