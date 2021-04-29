package variables

import "encoding/xml"

type Xlif struct {
	XMLName xml.Name `xml:"xliff"`
	Version string   `xml:"version,attr"`
	File    File     `xml:"file"`
}

type File struct {
	XMLName  xml.Name `xml:"file"`
	Original string   `xml:"original,attr"`
	Source   string   `xml:"source-language,attr"`
	Target   string   `xml:"target-language,attr"`
	Type     string   `xml:"translation-type,attr"`
	Datatype string   `xml:"datatype,attr"`
	Body     Body     `xml:"body"`
}

type Body struct {
	XMLName xml.Name `xml:"body"`
	Trans   []Trans  `xml:"trans-unit"`
}

type Trans struct {
	XMLName  xml.Name `xml:"trans-unit"`
	ID       string   `xml:"id,attr"`
	MaxWidth string   `xml:"maxwidth,attr"`
	SizeUnit string   `xml:"size-unit,attr"`
	Source   string   `xml:"source"`
	Target   string   `xml:"target"`
}
