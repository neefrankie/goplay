package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var httpClient = &http.Client{}

type Fetch struct {
	method    string
	url       string
	query     url.Values
	body      io.Reader
	header    http.Header
	errors    []error
	basicAuth BasicAuth
}

func New() *Fetch {
	return &Fetch{
		method:    "GET",
		url:       "",
		query:     url.Values{},
		body:      nil,
		header:    http.Header{},
		errors:    nil,
		basicAuth: BasicAuth{},
	}
}

func (f *Fetch) Get(url string) *Fetch {
	f.method = "GET"
	f.url = url

	return f
}

func (f *Fetch) Post(url string) *Fetch {
	f.method = "POST"
	f.url = url

	return f
}

func (f *Fetch) Put(url string) *Fetch {
	f.method = "PUT"
	f.url = url

	return f
}

func (f *Fetch) Patch(url string) *Fetch {
	f.method = "PATCH"
	f.url = url

	return f
}

func (f *Fetch) Delete(url string) *Fetch {
	f.method = "DELETE"
	f.url = url

	return f
}

func (f *Fetch) WithQuery(q url.Values) *Fetch {
	f.query = q
	return f
}

func (f *Fetch) SetQuery(key, value string) *Fetch {
	f.query.Set(key, value)

	return f
}

func (f *Fetch) SetQueryN(kv map[string]string) *Fetch {
	for k, v := range kv {
		f.query.Set(k, v)
	}

	return f
}

func (f *Fetch) AddQuery(key, value string) *Fetch {
	f.query.Add(key, value)
	return f
}

func (f *Fetch) AddQueryN(kv map[string]string) *Fetch {
	for k, v := range kv {
		f.query.Add(k, v)
	}

	return f
}

func (f *Fetch) WithHeader(h http.Header) *Fetch {
	f.header = h

	return f
}

func (f *Fetch) SetHeader(k, v string) *Fetch {
	f.header.Set(k, v)

	return f
}

func (f *Fetch) SetHeaderN(kv map[string]string) *Fetch {
	for k, v := range kv {
		f.header.Set(k, v)
	}

	return f
}

func (f *Fetch) AddHeader(key, value string) *Fetch {
	f.header.Add(key, value)
	return f
}

func (f *Fetch) AddHeaderN(kv map[string]string) *Fetch {
	for k, v := range kv {
		f.header.Add(k, v)
	}

	return f
}

func (f *Fetch) AcceptLang(v string) *Fetch {
	f.header.Set("Accept-Language", v)

	return f
}

func (f *Fetch) SetBearerAuth(key string) *Fetch {
	f.header.Set("Authorization", "Bearer "+key)

	return f
}

func (f *Fetch) SetBasicAuth(username, password string) *Fetch {
	f.basicAuth = BasicAuth{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	return f
}

func (f *Fetch) SetJSON() *Fetch {
	f.header.Add("Content-Type", ContentJSON)

	return f
}

func (f *Fetch) Stream(body io.Reader) *Fetch {
	f.body = body
	return f
}

func (f *Fetch) StreamJSON(body io.Reader) *Fetch {
	f.header.Add("Content-Type", ContentJSON)
	f.body = body

	return f
}

func (f *Fetch) SendJSONBlob(b []byte) *Fetch {
	return f.StreamJSON(bytes.NewReader(b))
}

func (f *Fetch) SendJSON(v interface{}) *Fetch {
	d, err := json.Marshal(v)
	if err != nil {
		f.errors = append(f.errors, err)

		return f
	}

	return f.StreamJSON(bytes.NewReader(d))
}

func (f *Fetch) End() (*http.Response, []error) {
	if f.errors != nil {
		return nil, f.errors
	}

	if len(f.query) != 0 {
		f.url = fmt.Sprintf("%s?%s", f.url, f.query.Encode())
	}

	req, err := http.NewRequest(f.method, f.url, f.body)
	if err != nil {
		f.errors = append(f.errors, err)
		return nil, f.errors
	}

	req.Header = f.header

	resp, err := httpClient.Do(req)
	if err != nil {
		f.errors = append(f.errors, err)
		return nil, f.errors
	}

	return resp, nil
}

// EndBlob reads response body and returns as a slice of bytes
func (f *Fetch) EndBlob() (Response, []error) {
	resp, errs := f.End()
	if errs != nil {
		return Response{}, f.errors
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		f.errors = append(f.errors, err)
		return Response{}, f.errors
	}

	return Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Body:       b,
	}, nil
}
