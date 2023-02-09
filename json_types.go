package central

import (
	"encoding/json"
	"time"

	"github.com/cydev/zero"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05 -07:00", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (d Date) IsEmpty() bool {
	return zero.IsZero(d)
}

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

func (d DateTime) IsEmpty() bool {
	return zero.IsZero(d)
}

type Time struct {
	time.Time
}

func (t Time) IsEmpty() bool {
	return t.Time.IsZero()
}

func (d Time) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

func (dt Time) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(dt.Time.Format("2006-01-02T15:04:05"))
}

func (dt *Time) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	dt.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	dt.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}
