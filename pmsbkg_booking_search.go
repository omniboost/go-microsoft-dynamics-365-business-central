package guestline

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewBookingSearchRequest() BookingSearchRequest {
	r := BookingSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookingSearchRequest struct {
	client      *Client
	queryParams *BookingSearchQueryParams
	pathParams  *BookingSearchPathParams
	method      string
	headers     http.Header
	requestBody BookingSearchRequestBody
}

func (r BookingSearchRequest) NewQueryParams() *BookingSearchQueryParams {
	return &BookingSearchQueryParams{}
}

type BookingSearchQueryParams struct {
}

func (p BookingSearchQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BookingSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r BookingSearchRequest) NewPathParams() *BookingSearchPathParams {
	return &BookingSearchPathParams{}
}

type BookingSearchPathParams struct {
}

func (p *BookingSearchPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BookingSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *BookingSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *BookingSearchRequest) Method() string {
	return r.method
}

func (r BookingSearchRequest) NewRequestBody() BookingSearchRequestBody {
	return BookingSearchRequestBody{}
}

type BookingSearchRequestBody struct {
	XMLName   xml.Name `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsbkg_BookingSearch"`
	SessionID string
	Filters   struct {
		BookRef string `xml:"BookRef"`
	} `xml:"Filters"`
}

func (r *BookingSearchRequest) RequestBody() *BookingSearchRequestBody {
	return &r.requestBody
}

func (r *BookingSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *BookingSearchRequest) SetRequestBody(body BookingSearchRequestBody) {
	r.requestBody = body
}

func (r *BookingSearchRequest) NewResponseBody() *BookingSearchResponseBody {
	return &BookingSearchResponseBody{}
}

type BookingSearchResponseBody struct {
	XMLName                   xml.Name       `xml:BookingSearchResponse`
	PmsbkgBookingSearchResult ExceptionBlock `xml:"pmsbkg_BookingSearchResult"`
	SearchResults             struct {
		Reservations []struct {
			RoomId        string `xml:"RoomId"`
			BookRef       string `xml:"BookRef"`
			RoomPickId    string `xml:"RoomPickId"`
			BookingType   string `xml:"BookingType"`
			BookingStatus string `xml:"BookingStatus"`
			RoomTypeCode  string `xml:"RoomTypeCode"`
			PackageCode   string `xml:"PackageCode"`
			Arrival       string `xml:"Arrival"`
			Departure     string `xml:"Departure"`
			Creation      string `xml:"Creation"`
			Guests        []struct {
				Name         string `xml:"Name"`
				Salutation   string `xml:"Salutation"`
				Forename     string `xml:"Forename"`
				Surname      string `xml:"Surname"`
				TypeOfPerson string `xml:"TypeOfPerson"`
				Gender       struct {
					Nil string `xml:"nil,attr"`
				} `xml:"Gender"`
				ProfileRef string `xml:"ProfileRef"`
				LoyaltyID  string `xml:"LoyaltyID"`
				FolioID    string `xml:"FolioID"`
			} `xml:"Guests>Guest"`
			PostsAllowed          string `xml:"PostsAllowed"`
			BookRefRoomPickID     string `xml:"BookRefRoomPickID"`
			AllowRoomMoves        string `xml:"AllowRoomMoves"`
			MarketSegment         string `xml:"MarketSegment"`
			TotalCostNett         string `xml:"TotalCostNett"`
			TotalCostGross        string `xml:"TotalCostGross"`
			LastEdited            string `xml:"LastEdited"`
			GDSRef                string `xml:"GDSRef"`
			CRSRef1               string `xml:"CRSRef1"`
			CRSRef2               string `xml:"CRSRef2"`
			SystemSource          string `xml:"SystemSource"`
			DistributionChannelId string `xml:"DistributionChannelId"`
			PreCheckIn            string `xml:"PreCheckIn"`
			RezlynxCRS            string `xml:"RezlynxCRS"`
		} `xml:"Reservations>Reservation"`
	} `xml:"SearchResults"`
}

func (r *BookingSearchRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx", r.PathParams())
	return &u
}

func (r *BookingSearchRequest) Do() (BookingSearchResponseBody, error) {
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
