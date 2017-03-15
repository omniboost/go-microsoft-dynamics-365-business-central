package multivers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-multivers/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
	apiPrefix      = "/api"
)

// Client manages communication with Unit4 Multivers API
type Client struct {
	// SOAP client used to communicate with the API.
	client *http.Client

	// Url pointing to base Unit4 Multivers API
	BaseURL *url.URL

	// Debugging flag
	Debug bool

	// User agent for client
	UserAgent string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// General data
	Administration          *AdministrationService
	AdministrationInfoList  *AdministrationInfoListService
	AdministrationGroupList *AdministrationGroupListService

	// Administration data
	ProductInfoList     *ProductInfoListService
	ProductTypeNVL      *ProductTypeNVLService
	ProjectEntryTypeNVL *ProjectEntryTypeNVLService
	VatCodeInfo         *VatCodeInfoService
	VatCodeInfoList     *VatCodeInfoListService
	ProductGroupNVL     *ProductGroupNVLService
	ProductGroup        *ProductGroupService
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// NewClient returns a new Unit4 Multivers API client
func NewClient(httpClient *http.Client, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:    httpClient,
		BaseURL:   nil,
		UserAgent: userAgent,
		Debug:     false,
	}

	c.SetBaseURL(baseURL)

	// General data
	c.Administration = NewAdministrationService(c)
	c.AdministrationInfoList = NewAdministrationInfoListService(c)
	c.AdministrationGroupList = NewAdministrationGroupListService(c)

	// Administration data
	c.ProductInfoList = NewProductInfoListService(c)
	c.ProductTypeNVL = NewProductTypeNVLService(c)
	c.ProjectEntryTypeNVL = NewProjectEntryTypeNVLService(c)
	c.VatCodeInfo = NewVatCodeInfoService(c)
	c.VatCodeInfoList = NewVatCodeInfoListService(c)
	c.ProductGroupNVL = NewProductGroupNVLService(c)
	c.ProductGroup = NewProductGroupService(c)

	// Commands

	// Documents

	return c
}

func (c *Client) SetDebug(debug bool) {
	c.Debug = debug
}

func (c *Client) SetSandbox(sandbox bool) {
	if sandbox == true {
		// u, _ := url.ParseRequestURI(companies.SandboxEndpoint)
		// c.Companies.Endpoint = u
	} else {
		// u, _ := url.ParseRequestURI(companies.Endpoint)
		// c.Companies.Endpoint = u
	}
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider(baseURL.String())

	// set base url for use in http client
	c.BaseURL = baseURL
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	u := c.GetEndpoint(path)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) GetEndpoint(path string) *url.URL {
	basePath := strings.TrimSuffix(c.BaseURL.Path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	u := *c.BaseURL
	u.Path = basePath + path
	return &u
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.Debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.Debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	err = json.NewDecoder(httpResp.Body).Decode(responseBody)
	return httpResp, err
}
