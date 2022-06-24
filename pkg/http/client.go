package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.elastic.co/apm/module/apmhttp"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	BaseUrl      string
	Headers      map[string]string
	DefaultQuery map[string]string
	HttpClient   *http.Client
	ErrorHandler func(res *http.Response, uri string) error
}

func (r *Request) SetTimeout(seconds time.Duration) {
	r.HttpClient.Timeout = time.Second * seconds
}

//goland:noinspection GoUnusedExportedFunction
func NewHttpClient(baseUrl string) Request {
	return Request{
		Headers:      make(map[string]string),
		HttpClient:   newDefaultClient(),
		DefaultQuery: make(map[string]string),
		ErrorHandler: DefaultErrorHandler,
		BaseUrl:      baseUrl,
	}
}

func NewHttpJSONClient(baseUrl string) Request {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return Request{
		Headers:      headers,
		HttpClient:   newDefaultClient(),
		ErrorHandler: DefaultErrorHandler,
		BaseUrl:      baseUrl,
	}
}

func newDefaultClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 200,
	}
}

var DefaultErrorHandler = func(res *http.Response, uri string) error {
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		return nil
	}
	return errors.New("error in making request")
}

//GetWithContext makes and HTTP GET call with the provided values after appending the default query values
//in addition to any existing ones.
func (r *Request) GetWithContext(result interface{}, path string, query url.Values, ctx context.Context) error {
	queryStr := r.safeGetQueryStrWithDefaults(query)
	uri := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result, ctx)
}

// Get makes an HTTP GET call after appending the default query values in addition to the existing ones.
func (r *Request) Get(result interface{}, path string, query url.Values) error {
	queryStr := r.safeGetQueryStrWithDefaults(query)
	uri := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result, context.Background())
}

//safeGetQueryStrWithDefaults appends the default queries (auth key and other specified data) without creating a panic.
//It returns the URL encoded query string ( "<url>?auth_key=value&item=value ..." )
func (r *Request) safeGetQueryStrWithDefaults(query url.Values) string {
	if query == nil {
		query = make(url.Values)
	}
	r.AppendDefaultQuery(&query)

	return query.Encode()
}

func (r *Request) Post(result interface{}, path string, body interface{}, query url.Values) error {
	buf, err := GetBody(body)
	if err != nil {
		return err
	}
	queryStr := r.safeGetQueryStrWithDefaults(query)
	uri := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	return r.Execute("POST", uri, buf, result, context.Background())
}

func (r *Request) PostWithContext(result interface{}, path string, body interface{}, ctx context.Context) error {
	buf, err := GetBody(body)
	if err != nil {
		return err
	}
	uri := r.GetBase(path)
	return r.Execute("POST", uri, buf, result, ctx)
}

func (r *Request) Execute(method string, url string, body io.Reader, result interface{}, ctx context.Context) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	c := apmhttp.WrapClient(r.HttpClient)
	res, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	err = r.ErrorHandler(res, url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if result == nil {
		return nil
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}
	return err
}

func (r *Request) GetBase(path string) string {
	if path == "" {
		return r.BaseUrl
	}
	return fmt.Sprintf("%s/%s", r.BaseUrl, path)
}

//AppendDefaultQuery add the specified default query values to a query
func (r *Request) AppendDefaultQuery(query *url.Values) {
	for k, v := range r.DefaultQuery {
		if existingValue := query.Get(k); len(existingValue) > 0 {
			continue
		}
		query.Add(k, v)
	}
}

func GetBody(body interface{}) (buf io.ReadWriter, err error) {
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
	}
	return
}

//QueryParamHelper accepts a map of string -> interface and returns the supported url.Values for the map.
func QueryParamHelper(queryParams map[string]interface{}) (urlVals url.Values) {
	if queryParams != nil {
		urlVals = make(url.Values)
		for key, val := range queryParams {
			switch val.(type) {
			default:
				urlVals.Add(key, fmt.Sprint(val))
			case []string:
				urlVals[key] = val.([]string)

			case string:
				urlVals.Add(key, val.(string))
			}
		}
	}
	return
}
