package dkplus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-dkplus/utils"
)

func (c *Client) NewCompanyGetRequest() CompanyGetRequest {
	r := CompanyGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompanyGetRequest struct {
	client      *Client
	queryParams *CompanyGetQueryParams
	pathParams  *CompanyGetPathParams
	method      string
	headers     http.Header
	requestBody CompanyGetRequestBody
}

func (r CompanyGetRequest) NewQueryParams() *CompanyGetQueryParams {
	return &CompanyGetQueryParams{}
}

type CompanyGetQueryParams struct {
}

func (p CompanyGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CompanyGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CompanyGetRequest) NewPathParams() *CompanyGetPathParams {
	return &CompanyGetPathParams{}
}

type CompanyGetPathParams struct {
}

func (p *CompanyGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompanyGetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CompanyGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompanyGetRequest) Method() string {
	return r.method
}

func (r CompanyGetRequest) NewRequestBody() CompanyGetRequestBody {
	return CompanyGetRequestBody{}
}

type CompanyGetRequestBody struct {
}

func (r *CompanyGetRequest) RequestBody() *CompanyGetRequestBody {
	return &r.requestBody
}

func (r *CompanyGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CompanyGetRequest) SetRequestBody(body CompanyGetRequestBody) {
	r.requestBody = body
}

func (r *CompanyGetRequest) NewResponseBody() *CompanyGetResponseBody {
	return &CompanyGetResponseBody{}
}

type CompanyGetResponseBody struct {
	Information struct {
		Owner     string `json:"Owner"`
		OwnerName string `json:"OwnerName"`
		License   string `json:"License"`
		Company   struct {
			Number       string `json:"Number"`
			Name         string `json:"Name"`
			SSNumber     string `json:"SSNumber"`
			Address1     string `json:"Address1"`
			Address2     string `json:"Address2"`
			Address3     string `json:"Address3"`
			ZipCode      string `json:"ZipCode"`
			Country      string `json:"Country"`
			BankCode     string `json:"BankCode"`
			BankAccGroup string `json:"BankAccGroup"`
			BankAccount  string `json:"BankAccount"`
			VATNumber    string `json:"VATNumber"`
			Phone        string `json:"Phone"`
			Mobile       string `json:"Mobile"`
			Fax          string `json:"Fax"`
			Email        string `json:"Email"`
			URL          string `json:"Url"`
			Swift        string `json:"Swift"`
			IBAN         string `json:"IBAN"`
			City         string `json:"City"`
		} `json:"Company"`
	} `json:"Information"`
	Customer struct {
		Enabled bool `json:"Enabled"`
	} `json:"Customer"`
	Product struct {
		Enabled   bool `json:"Enabled"`
		Warehouse struct {
			Enabled bool   `json:"Enabled"`
			Default string `json:"Default"`
		} `json:"Warehouse"`
		Categories bool `json:"Categories"`
	} `json:"Product"`
	Vendor struct {
		Enabled      bool `json:"Enabled"`
		Confirmation bool `json:"Confirmation"`
	} `json:"Vendor"`
	Sale struct {
		Enabled bool `json:"Enabled"`
	} `json:"Sale"`
	Project struct {
		Enabled             bool `json:"Enabled"`
		PhaseEnabled        bool `json:"PhaseEnabled"`
		TaskEnabled         bool `json:"TaskEnabled"`
		EnterDriveInJournal bool `json:"EnterDriveInJournal"`
		DefaultDriveUnit    int  `json:"DefaultDriveUnit"`
	} `json:"Project"`
	Dimmension struct {
		Dim1Enabled bool `json:"Dim1Enabled"`
		Dim2Enabled bool `json:"Dim2Enabled"`
		Dim3Enabled bool `json:"Dim3Enabled"`
	} `json:"Dimmension"`
	General struct {
		CurrencyEnabled bool   `json:"CurrencyEnabled"`
		DefaultCurrency string `json:"DefaultCurrency"`
		Attachments     bool   `json:"Attachments"`
	} `json:"General"`
	Member struct {
		Enabled bool `json:"Enabled"`
	} `json:"Member"`
}

func (r *CompanyGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("company", r.PathParams())
	return &u
}

func (r *CompanyGetRequest) Do() (CompanyGetResponseBody, error) {
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
