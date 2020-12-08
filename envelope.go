package guestline

import (
	"encoding/xml"
)

type RequestEnvelope struct {
	XMLName    xml.Name
	Namespaces []xml.Attr `xml:"-"`

	Header Header `xml:"soap:Header"`
	Body   Body   `xml:"soap:Body"`
}

type ResponseEnvelope struct {
	XMLName    xml.Name
	Namespaces []xml.Attr `xml:"-"`

	Header Header `xml:"Header"`
	Body   Body   `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "soap:Envelope"}

	for _, ns := range env.Namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type Body struct {
	Action     string      `xml:"-"`
	ActionBody interface{} `xml:",any"`
}

type Header interface{}

type ActionBody interface{}
