package models

import (
	"encoding/xml"

	"github.com/labstack/echo/v4"
)

type Action func(c echo.Context, envelope Envelope) error

type Envelope struct {
	Namespaces map[string]string `xml:"-"`
	Header     Header
	Body       Body
}

type Body struct {
  XMLName xml.Name `xml:"soapenv:Body"`
	Content interface{}
	Fault   Fault `xml:"Fault"`
}

type Header struct {
  XMLName xml.Name   `xml:"soapenv:Header"`
	Content string     `xml:",chardata"`
	Attrs   []xml.Attr `xml:",any,attr"`
}

type Fault struct {
	Code   string `xml:"faultcode"`
	String string `xml:"faultstring"`
	Actor  string `xml:"faultactor"`
	Detail Detail `xml:"detail"`
}

type Detail struct {
	XMLName xml.Name
	Content string `xml:",chardata"`
}
