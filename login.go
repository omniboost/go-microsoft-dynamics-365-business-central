package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-guestline/utils"
)

func (c *Client) NewLoginRequest() LoginRequest {
	r := LoginRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LoginRequest struct {
	client      *Client
	queryParams *LoginQueryParams
	pathParams  *LoginPathParams
	method      string
	headers     http.Header
	requestBody LoginRequestBody
}

func (r LoginRequest) NewQueryParams() *LoginQueryParams {
	return &LoginQueryParams{}
}

type LoginQueryParams struct {
}

func (p LoginQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LoginRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r LoginRequest) NewPathParams() *LoginPathParams {
	return &LoginPathParams{}
}

type LoginPathParams struct {
}

func (p *LoginPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LoginRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *LoginRequest) SetMethod(method string) {
	r.method = method
}

func (r *LoginRequest) Method() string {
	return r.method
}

func (r LoginRequest) NewRequestBody() LoginRequestBody {
	return LoginRequestBody{
		SiteID:       r.client.SiteID(),
		InterfaceID:  r.client.InterfaceID(),
		OperatorCode: r.client.OperatorCode(),
		Password:     r.client.Password(),
	}
}

type LoginRequestBody struct {
	XMLName      xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 LogIn"`
	SiteID       string
	InterfaceID  string
	OperatorCode string
	Password     string
}

func (r *LoginRequest) RequestBody() *LoginRequestBody {
	return &r.requestBody
}

func (r *LoginRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *LoginRequest) SetRequestBody(body LoginRequestBody) {
	r.requestBody = body
}

func (r *LoginRequest) NewResponseBody() *LoginResponseBody {
	return &LoginResponseBody{}
}

type LoginResponseBody struct {
	XMLName     xml.Name       `xml:LogInResponse`
	LogInResult ExceptionBlock `xml:"LogInResult"`
	SessionID   string         `xml:"SessionID"`
}

func (r *LoginRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *LoginRequest) Do() (LoginResponseBody, error) {
	// if r.QueryParams().ConsumerToken == "" {
	// 	r.QueryParams().ConsumerToken = r.client.ConsumerToken()
	// }

	// if r.QueryParams().EmployeeToken == "" {
	// 	r.QueryParams().EmployeeToken = r.client.EmployeeToken()
	// }

	// if r.QueryParams().ExpirationDate.IsZero() {
	// 	r.QueryParams().ExpirationDate = Date{time.Now().AddDate(0, 0, 1)}
	// }

	// fetch a new token if it isn't set already
	// if r.client.token == "" {
	// 	var err error
	// 	r.client.token, err = r.client.NewToken()
	// 	if err != nil {
	// 		return *r.NewResponseBody(), err
	// 	}
	// }

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
	return *responseBody, err
}
