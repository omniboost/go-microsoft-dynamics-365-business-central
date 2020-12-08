package guestline

import "time"

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

type DateTime struct {
	time.Time
}
