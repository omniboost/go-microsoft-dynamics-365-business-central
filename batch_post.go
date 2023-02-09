package central

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/omniboost/go-microsoft-dynamics-365-business-central/utils"
)

func (c *Client) NewBatchPost() BatchPost {
	r := BatchPost{
		client: c,
		method: http.MethodPost,
		headers: http.Header{
			"Isolation": []string{"snapshot"},
		},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BatchPost struct {
	client      *Client
	queryParams *BatchPostQueryParams
	pathParams  *BatchPostPathParams
	method      string
	headers     http.Header
	requestBody BatchPostBody
}

func (r BatchPost) NewQueryParams() *BatchPostQueryParams {
	return &BatchPostQueryParams{}
}

type BatchPostQueryParams struct {
}

func (p BatchPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BatchPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r BatchPost) NewPathParams() *BatchPostPathParams {
	return &BatchPostPathParams{}
}

type BatchPostPathParams struct {
	EnvironmentName string `schema:"environmentName"`
	CompanyID       string `schema:"companyID"`
	JournalID       string `schema:"journalID"`
}

func (p *BatchPostPathParams) Params() map[string]string {
	return map[string]string{
		"environmentName": p.EnvironmentName,
		"companyID":       p.CompanyID,
		"journalID":       p.JournalID,
	}
}

func (r *BatchPost) PathParams() *BatchPostPathParams {
	return r.pathParams
}

func (r *BatchPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BatchPost) SetMethod(method string) {
	r.method = method
}

func (r *BatchPost) Method() string {
	return r.method
}

func (r *BatchPost) Headers() http.Header {
	return r.headers
}

func (r BatchPost) NewRequestBody() BatchPostBody {
	return BatchPostBody{}
}

type BatchPostBody struct {
	Requests []BatchPostBodyRequest `json:"requests"`
}

type BatchPostBodyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body,omitempty"`
}

func NewBatchPostBodyRequest() BatchPostBodyRequest {
	return BatchPostBodyRequest{
		Headers: map[string]string{},
	}
}

func (r *BatchPost) RequestBody() *BatchPostBody {
	return &r.requestBody
}

func (r *BatchPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *BatchPost) SetRequestBody(body BatchPostBody) {
	r.requestBody = body
}

func (r *BatchPost) NewResponseBody() *BatchPostResponseBody {
	return &BatchPostResponseBody{}
}

type BatchPostResponseBody struct {
	Responses []BatchPostBodyResponse `json:"responses"`
}

type BatchPostBodyResponse struct {
	ID      interface{}       `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    json.RawMessage   `json:"body"`
}

func (r BatchPostBodyResponse) ToHTTPResponse() (*http.Response, error) {
	resp := &http.Response{
		Header: http.Header{},
	}

	resp.StatusCode = r.Status

	for k, v := range r.Headers {
		resp.Header.Add(k, v)
	}

	resp.ContentLength = -1
	resp.Body = io.NopCloser(bytes.NewReader(r.Body))
	return resp, nil
}

func (r *BatchPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2.0/{{.environmentName}}/api/v2.0/$batch", r.PathParams())
	// u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *BatchPost) Do() (BatchPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, err
	}

	// do error check on the sub responses
	var errors error
	for _, resp := range responseBody.Responses {
		httpResp, err := resp.ToHTTPResponse()
		if err != nil {
			return *responseBody, err
		}

		err = r.client.CheckResponse(httpResp)
		if err != nil {
			errors = multierror.Append(errors, err)
		}
	}

	if errors != nil {
		return *responseBody, errors
	}

	return *responseBody, nil
}

// func (r *BatchPost) FromHTTPRequest(req *http.Request) (BatchPostBodyRequest, error) {
// 	var err error
// 	body := NewBatchPostBodyRequest()

// 	// body.Headers["Company"] = "CRONUS NL"
// 	for k, vv := range req.Header {
// 		for _, v := range vv {
// 			body.Headers[k] = v
// 		}
// 	}

// 	body.Method = req.Method
// 	base := strings.TrimSuffix(r.URL().String(), "$batch")
// 	body.URL = strings.TrimPrefix(req.URL.String(), base)
// 	body.Body, err = ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		return body, err
// 	}

// 	return body, nil
// }

func (r *BatchPost) FromRequest(req Request) (BatchPostBodyRequest, error) {
	var err error
	body := NewBatchPostBodyRequest()

	// body.Headers["Company"] = "CRONUS NL"
	for k, vv := range req.Headers() {
		for _, v := range vv {
			body.Headers[k] = v
		}
	}

	body.Method = req.Method()
	base := strings.TrimSuffix(r.URL().String(), "$batch")
	body.URL = strings.TrimPrefix(req.URL().String(), base)

	body.Body, err = json.Marshal(req.RequestBodyInterface())
	if err != nil {
		return body, err
	}

	return body, nil
}
