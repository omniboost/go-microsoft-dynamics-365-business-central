package multivers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		// status code is ok: no errors
		return nil
	}

	// create base error response
	errorResponse := &ErrorResponse{Response: r, Code: r.StatusCode}

	// check if content type is json
	err := checkContentType(r)
	if err != nil {
		errorResponse.Message = err.Error()
		return errorResponse
	}

	// read response body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errorResponse
	}

	// no data returned: error
	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Message = fmt.Sprintf("Malformed json response")
		return errorResponse
	}

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// HTTP status code
	Code int

	// Fault message
	Message string `json:"Message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d (%v)",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != "application/json" {
		return fmt.Errorf("Expected Content-Type \"application/json\", got \"%s\"", contentType)
	}

	return nil
}
