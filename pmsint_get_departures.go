package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewGetDeparturesRequest() GetDeparturesRequest {
	r := GetDeparturesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetDeparturesRequest struct {
	client      *Client
	queryParams *GetDeparturesQueryParams
	pathParams  *GetDeparturesPathParams
	method      string
	headers     http.Header
	requestBody GetDeparturesRequestBody
}

func (r GetDeparturesRequest) NewQueryParams() *GetDeparturesQueryParams {
	return &GetDeparturesQueryParams{}
}

type GetDeparturesQueryParams struct {
}

func (p GetDeparturesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDeparturesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetDeparturesRequest) NewPathParams() *GetDeparturesPathParams {
	return &GetDeparturesPathParams{}
}

type GetDeparturesPathParams struct {
}

func (p *GetDeparturesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDeparturesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetDeparturesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDeparturesRequest) Method() string {
	return r.method
}

func (r GetDeparturesRequest) NewRequestBody() GetDeparturesRequestBody {
	return GetDeparturesRequestBody{}
}

type GetDeparturesRequestBody struct {
	XMLName       xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetDepartures"`
	SessionID     string
	DepartureDate DateTime
}

func (r *GetDeparturesRequest) RequestBody() *GetDeparturesRequestBody {
	return &r.requestBody
}

func (r *GetDeparturesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetDeparturesRequest) SetRequestBody(body GetDeparturesRequestBody) {
	r.requestBody = body
}

func (r *GetDeparturesRequest) NewResponseBody() *GetDeparturesResponseBody {
	return &GetDeparturesResponseBody{}
}

type GetDeparturesResponseBody struct {
	XMLName                   xml.Name       `xml:GetDeparturesResponse`
	PmsintGetDeparturesResult ExceptionBlock `xml:"pmsint_GetDeparturesResult"`
	Departures                Departures     `xml:"GetDepartures>Departures>cpmsint_GetDepartures_DepartureItem"`
}

func (r *GetDeparturesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetDeparturesRequest) Do() (GetDeparturesResponseBody, error) {
	var err error

	// fetch a new token if it isn't set already
	r.requestBody.SessionID, err = r.client.SessionID()
	if err != nil {
		return *r.NewResponseBody(), err
	}

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

type Departures []Departure

type Departure struct {
	DepositPaid    string `xml:"DepositPaid"`
	DepositDue     string `xml:"DepositDue"`
	Infants        string `xml:"Infants"`
	Children       string `xml:"Children"`
	Adults         string `xml:"Adults"`
	RoomType       string `xml:"RoomType"`
	Package        string `xml:"Package"`
	Company        string `xml:"Company"`
	Notes          string `xml:"Notes"`
	BookRef        string `xml:"BookRef"`
	BookRefRoomRef string `xml:"BookRefRoomRef"`
	ETA            string `xml:"ETA"`
	Salutation     string `xml:"Salutation"`
	Surname        string `xml:"Surname"`
	Forename       string `xml:"Forename"`
	RoomID         string `xml:"RoomID"`
	ETD            string `xml:"ETD"`
	PrivateNotes   string `xml:"PrivateNotes"`
	PublicNotes    string `xml:"PublicNotes"`
	CustomNotes1   string `xml:"CustomNotes1"`
	CustomNotes2   string `xml:"CustomNotes2"`
	CustomNotes3   string `xml:"CustomNotes3"`
	ExternalNotes  string `xml:"ExternalNotes"`
	RoomBalance    string `xml:"RoomBalance"`
	Master         string `xml:"Master"`
}
