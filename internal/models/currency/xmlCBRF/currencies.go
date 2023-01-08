package xmlCBRF

import "encoding/xml"

type Currencies struct {
	XMLName  xml.Name   `xml:"ValCurs"`
	Currency []Currency `xml:"Valute"`
}
