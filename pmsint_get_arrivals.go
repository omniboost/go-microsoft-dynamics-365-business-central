package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewGetArrivalsRequest() GetArrivalsRequest {
	r := GetArrivalsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetArrivalsRequest struct {
	client      *Client
	queryParams *GetArrivalsQueryParams
	pathParams  *GetArrivalsPathParams
	method      string
	headers     http.Header
	requestBody GetArrivalsRequestBody
}

func (r GetArrivalsRequest) NewQueryParams() *GetArrivalsQueryParams {
	return &GetArrivalsQueryParams{}
}

type GetArrivalsQueryParams struct {
}

func (p GetArrivalsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetArrivalsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetArrivalsRequest) NewPathParams() *GetArrivalsPathParams {
	return &GetArrivalsPathParams{}
}

type GetArrivalsPathParams struct {
}

func (p *GetArrivalsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetArrivalsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetArrivalsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetArrivalsRequest) Method() string {
	return r.method
}

func (r GetArrivalsRequest) NewRequestBody() GetArrivalsRequestBody {
	return GetArrivalsRequestBody{}
}

type GetArrivalsRequestBody struct {
	XMLName     xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsint_GetArrivals"`
	SessionID   string
	ArrivalDate DateTime
}

func (r *GetArrivalsRequest) RequestBody() *GetArrivalsRequestBody {
	return &r.requestBody
}

func (r *GetArrivalsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetArrivalsRequest) SetRequestBody(body GetArrivalsRequestBody) {
	r.requestBody = body
}

func (r *GetArrivalsRequest) NewResponseBody() *GetArrivalsResponseBody {
	return &GetArrivalsResponseBody{}
}

type GetArrivalsResponseBody struct {
	XMLName                 xml.Name       `xml:GetArrivalsResponse`
	PmsintGetArrivalsResult ExceptionBlock `xml:"pmsint_GetArrivalsResult"`
	Arrivals                Arrivals       `xml:"GetArrivals>Arrivals>cpmsint_GetArrivals_ArrivalItem"`
}

func (r *GetArrivalsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *GetArrivalsRequest) Do() (GetArrivalsResponseBody, error) {
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

type Arrivals []Arrival

type Arrival struct {
	DepositPaid              string `xml:"DepositPaid"`
	DepositDue               string `xml:"DepositDue"`
	Infants                  string `xml:"Infants"`
	Children                 string `xml:"Children"`
	Adults                   string `xml:"Adults"`
	RoomType                 string `xml:"RoomType"`
	Package                  string `xml:"Package"`
	Company                  string `xml:"Company"`
	Notes                    string `xml:"Notes"`
	BookRef                  string `xml:"BookRef"`
	BookRefRoomRef           string `xml:"BookRefRoomRef"`
	ETA                      string `xml:"ETA"`
	Salutation               string `xml:"Salutation"`
	Surname                  string `xml:"Surname"`
	Forename                 string `xml:"Forename"`
	RoomID                   string `xml:"RoomID"`
	ProfileReference         string `xml:"ProfileReference"`
	DepositOutstanding       string `xml:"DepositOutstanding"`
	TotalNights              string `xml:"TotalNights"`
	PrivateNotes             string `xml:"PrivateNotes"`
	PublicNotes              string `xml:"PublicNotes"`
	CustomNotes1             string `xml:"CustomNotes1"`
	CustomNotes2             string `xml:"CustomNotes2"`
	CustomNotes3             string `xml:"CustomNotes3"`
	ExternalNotes            string `xml:"ExternalNotes"`
	PreCalcChargesTotalGross string `xml:"PreCalcChargesTotalGross"`
	PreCalcChargesTotalNett  string `xml:"PreCalcChargesTotalNett"`
	Master                   string `xml:"Master"`
}
