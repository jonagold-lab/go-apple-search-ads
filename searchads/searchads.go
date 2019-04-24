package searchads

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.searchads.apple.com/api/v2/"
	userAgent      = "go-apple-search-ads"
)

// A Client manages communication with the Google Search Ads API.
type Client struct {
	client                  *http.Client // Reuse a single struct instead of allocating one for each service on the heap.
	BaseURL                 *url.URL
	OrgID                   *int64
	UserAgent               string
	common                  service
	Campaign                *CampaignService
	AdGroup                 *AdGroupService
	ACL                     *ACLService
	CampaignNegativeKeyword *CampaignNegativeKeywordServive
	AdGroupNegativeKeyword  *AdGroupNegativeKeywordServive
	AdGroupTargetingKeyword *AdGroupTargetingKeywordServive
	Report                  *ReportService
}

type service struct {
	client *Client
}

// ListOptions to hold url params like pagination
type ListOptions struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func setClientWithCerts(pemFile, keyFile string) (*http.Client, error) {
	cert, err := tls.LoadX509KeyPair(pemFile, keyFile)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}, nil
}

// NewClient with either http.Client or with pemFile and keyFile
func NewClient(httpClient *http.Client, pemFile, keyFile string, orgID *int64) (*Client, error) {
	if httpClient == nil {
		var err error
		httpClient, err = setClientWithCerts(pemFile, keyFile)
		if err != nil {
			return nil, err
		}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, OrgID: orgID}
	c.common.client = c
	c.Campaign = (*CampaignService)(&c.common)
	c.CampaignNegativeKeyword = (*CampaignNegativeKeywordServive)(&c.common)
	c.AdGroup = (*AdGroupService)(&c.common)
	c.AdGroupNegativeKeyword = (*AdGroupNegativeKeywordServive)(&c.common)
	c.AdGroupTargetingKeyword = (*AdGroupTargetingKeywordServive)(&c.common)
	c.ACL = (*ACLService)(&c.common)
	c.Report = (*ReportService)(&c.common)
	return c, nil
}

// NewRequest to build request
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if c.OrgID != nil {
		req.Header.Set("Authorization", fmt.Sprintf("orgId=%v", *c.OrgID))
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do to execute request and handle repsonse of all services
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		// If context is cancled then return that error
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		// If the error type is *url.Error
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()
	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}
	type Response struct {
		Data       interface{}
		Pagination Pagination
	}
	rv := &Response{Data: v}
	err = json.NewDecoder(resp.Body).Decode(rv)
	response.Pagination = rv.Pagination
	return response, err
}

// Response to hold apple serach ads response with page details
type Response struct {
	*http.Response
	Pagination
}

// Pagination Struct to hold pagination information
type Pagination struct {
	TotalResults int `json:"totalResults"`
	StartIndex   int `json:"startIndex"`
	ItemsPerPage int `json:"itemsPerPage"`
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// CheckResponse to build an erro if code is not 2xx
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

/*
An ErrorResponse reports one or more errors caused by an API request.
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Errors   Errors         `json:"error"` // more detail on individual errors
}

// Errors struct holds all messages
type Errors struct {
	Messages []ErrorMessage `json:"errors"`
}

// ErrorMessage with details
type ErrorMessage struct {
	MessageCode string `json:"messageCode"`
	Message     string `json:"message"`
	Field       string `json:"field"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Errors)
}
