package guestline

type ExceptionBlock struct {
	ExceptionCode        string `xml:"ExceptionCode"`
	ExceptionDescription string `xml:"ExceptionDescription"`
}
