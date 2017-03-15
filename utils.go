package multivers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aodin/date"
	"github.com/google/go-querystring/query"
)

func addQueryParamsToRequest(requestParams interface{}, req *http.Request, skipEmpty bool) error {
	params, err := query.Values(requestParams)
	if err != nil {
		return err
	}

	query := req.URL.Query()

	for k, vals := range params {
		for _, v := range vals {
			if skipEmpty && v == "" {
				continue
			}

			query.Add(k, v)
		}
	}

	req.URL.RawQuery = query.Encode()
	return nil
}

type DateNLNL struct {
	date.Date
}

// func (d *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	t := time.Time(*d)
// 	return e.EncodeElement(t.Format("20060102"), start)
// }

func (d *DateNLNL) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// 28-1-2008
	layout := "2-1-2006"
	time, err := time.Parse(layout, value)
	d.Date = date.FromTime(time)
	return err
}
