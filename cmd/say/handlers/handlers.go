package handlers

import (
	"CPaaS/pkg/domains/say"
	"encoding/xml"
	"net/http"
)

const Header = `<?xml version="1.0"?>` + "\n"

func SaySomethingHandler(w http.ResponseWriter, r *http.Request) {
	inbound := &say.ResponseSay{
		Pause: say.Pause{
			Length: 5,
		},
		Say: say.Say{
			Value:    "How are you today",
			Voice:    "woman",
			Language: "en-US",
			Loop:     3,
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
