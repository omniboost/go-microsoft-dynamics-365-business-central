package multivers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/aodin/date"
	"github.com/gorilla/schema"
)

func addQueryParamsToRequest(requestParams interface{}, req *http.Request, skipEmpty bool) error {
	params := url.Values{}
	encoder := schema.NewEncoder()
	err := encoder.Encode(requestParams, params)
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

func NewDateNLNL(year int, month time.Month, day int) DateNLNL {
	d := date.New(year, month, day)
	return DateNLNL{Date: d}
}

// func (d *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	t := time.Time(*d)
// 	return e.EncodeElement(t.Format("20060102"), start)
// }

func (d *DateNLNL) MarshalJSON() ([]byte, error) {
	layout := "2-1-2006"
	fmtDate := d.Date.Format(layout)
	return json.Marshal(fmtDate)
}

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
