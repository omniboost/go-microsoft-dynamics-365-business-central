package central

import (
	"net/http"
	"net/url"
)

type Request interface {
	Method() string
	Headers() http.Header
	// QueryParams() QueryParams
	PathParamsInterface() PathParams
	RequestBodyInterface() interface{}
	URL() *url.URL
}

type QueryParams interface {
	ToURLValues() (url.Values, error)
}

type PathParams interface {
	Params() map[string]string
}
