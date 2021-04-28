package variables

import "encoding/xml"

type Xlif struct {
	XMLName xml.Name `xml:"xliff"`
	File    File     `xml:"file"`
}

type File struct {
	XMLName xml.Name `xml:"file"`
	Body    Body     `xml:"body"`
}

type Body struct {
	XMLName xml.Name `xml:"body"`
	Trans   []Trans  `xml:"trans-unit"`
}

type Trans struct {
	XMLName xml.Name `xml:"trans-unit"`
	Source  string   `xml:"source"`
	Target  string   `xml:"target"`
}
