package main

import (
	c "CPaaS/config"
	s "CPaaS/services"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	config := c.NewConfig()
	cpaas := s.NewClient(config)
	cpaas.Authentication.Authenticate(s.On)
	if cpaas.Authentication.IsAuthenticated() {
		//boughtNumber := cpaas.Number.ReadBuyNumberResponse(config.ResultBuyNumber)
		//fmt.Printf("The number %s was last updated in %s.\n", strings.TrimSpace(boughtNumber.FriendlyName), strings.TrimSpace(boughtNumber.DateUpdated))
		//phoneNumber, _ := cpaas.Number.BuyNumber(s.Bash, s.AC647, s.Off)
		//if !phoneNumber.IsBought {
		var numbers []s.IncomingPhoneNumer
		numberAPIResponse, _ := cpaas.Number.ListNumbers(s.On)
		numbers = numberAPIResponse.IncomingPhoneNumbers
		phoneNumber := s.PickRandonNumber(numbers)
		fmt.Printf("The number %s was purchased in %s.\n", strings.TrimSpace(phoneNumber.FriendlyName), strings.TrimSpace(phoneNumber.DateUpdated))
		//}
	}
	r := mux.NewRouter()
	http.Handle("/", r)
	addr := ":5000"

	r.HandleFunc("/PLAY1", PlayFirst).Methods("GET")
	r.HandleFunc("/PLAY2", PlaySecond).Methods("GET")
	r.HandleFunc("/SAY", Say).Methods("GET")
	r.HandleFunc("/DTMF1", DTMF1).Methods("GET")
	r.HandleFunc("/DTMF2", DTMF2).Methods("GET")
	r.HandleFunc("/DTMF3", DTMF3).Methods("GET")
	r.HandleFunc("/Dial", Dial).Methods("GET")
	r.HandleFunc("/SMS", SMS).Methods("GET")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err)
		http.ListenAndServe(addr, nil)
	}
}

func PlayFirst(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponsePlay{
		Play: s.Play{
			Value: "https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/Coldplay/Paradise.mp3",
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/xml")
	w.Write(iXML)
	return
}

func PlaySecond(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponsePlay{
		Play: s.Play{
			Value:  "https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/Alok/Hear_Me_Now.mp3",
			Method: "POST",
			Loop:   5,
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/xml")
	w.Write(iXML)
	return
}

func Dial(w http.ResponseWriter, r *http.Request) {
	inbound := &s.Response{
		Dial: s.Dial{
			Value:  "(647) 695-6429",
			Method: "POST",
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(iXML)
	return
}

func Say(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseSay{
		Say: s.Say{
			Value:    "How are you?",
			Voice:    "woman",
			Language: "fr-CA",
			Loop:     0,
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(iXML)
	return
}

func DTMF1(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseGather{
		Gather: s.GatherSay{
			Method:    "POST",
			NumDigits: 1,
			Input:     "dtmf",
			Say: s.Say{
				Value:    "Hello",
				Voice:    "woman",
				Language: "en-us",
				Loop:     0,
			},
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(iXML)
	return
}

func DTMF2(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseGather{
		Gather: s.GatherSay{
			Method:    "POST",
			NumDigits: 1,
			Input:     "dtmf",
			Say: s.Say{
				Value:    "How are you?",
				Voice:    "woman",
				Language: "en-us",
				Loop:     0,
			},
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(iXML)
	return
}
func DTMF3(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseGather{
		Gather: s.GatherSay{
			Method:    "POST",
			NumDigits: 1,
			Input:     "dtmf",
			Say: s.Say{
				Value:    "Everything is fine",
				Voice:    "woman",
				Language: "en-us",
				Loop:     0,
			},
		},
	}
	iXML, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(iXML)
	return
}
func SMS(w http.ResponseWriter, r *http.Request) {
	return
}
