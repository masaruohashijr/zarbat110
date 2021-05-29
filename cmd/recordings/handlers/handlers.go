package handlers

import (
	"CPaaS/pkg/domains/records"
	"CPaaS/pkg/domains/say"
	"encoding/xml"
	"net/http"
)

func RecordingHandler(w http.ResponseWriter, r *http.Request) {
	println("RecordingHandler")
	inbound := &records.ResponseRecord{
		Say: say.Say{
			Value: "Hi. Starting to record this conversation.",
		},
		Record: records.Record{
			Action:             "",
			Background:         false,
			Method:             "",
			TimeOut:            30,
			MaxLength:          60,
			Transcribe:         true,
			TranscribeCallback: "https://43a0103ec72e.ngrok.io/ReceiveTranscription",
			TranscribeQuality:  "auto",
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	println("*************************************")
	xml := string(iXML)
	println(xml)
	w.Write([]byte(xml))
}
