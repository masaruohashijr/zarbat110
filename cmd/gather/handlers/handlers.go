package handlers

import (
	"CPaaS/pkg/domains/gather"
	"CPaaS/pkg/domains/say"
	"encoding/xml"
	"net/http"
)

const Header = `<?xml version="1.0"?>` + "\n"

func GatherhHandler(w http.ResponseWriter, r *http.Request) {
	inbound := &gather.ResponseGather{
		Pause: say.Pause{
			Length: 5,
		},
		Gather: gather.Gather{
			Input:    "speech",
			Language: "en-US",
			Timeout:  15,
			Action:   "https://50d980b2d2bf.ngrok.io/SpeechResult",
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	println("*************************************")
	xml := Header + string(iXML)
	println(xml)
	w.Write([]byte(xml))
}

func SpeechResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sr := r.FormValue("SpeechResult")
	println("************************************************")
	println("SpeechResult")
	println(sr)
	println("Done")
}
