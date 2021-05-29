package records

import (
	"CPaaS/pkg/domains/say"
	"encoding/xml"
)

type ResponseRecord struct {
	XMLName xml.Name `xml:"Response"`
	Say     say.Say  `xml:"Say"`
	Record  Record   `xml:"Record"`
}

type Record struct {
	Action             string `xml:"action,attr,omitempty"`
	Background         bool   `xml:"background,attr,omitempty"`
	Method             string `xml:"method,attr,omitempty"`
	TimeOut            int    `xml:"timeout,attr,omitempty"`
	MaxLength          int    `xml:"maxlength,attr,omitempty"`
	Transcribe         bool   `xml:"transcribe,attr,omitempty"`
	TranscribeCallback string `xml:"transcribeCallback,attr,omitempty"`
	TranscribeQuality  string `xml:"transcribeQuality,attr,omitempty"`
}

type Pause struct {
	Length int `xml:"length,attr"`
}
