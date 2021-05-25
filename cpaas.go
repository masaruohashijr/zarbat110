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
	addr := ":5005"

	r.HandleFunc("/PLAY1", PlayFirst).Methods("POST")
	r.HandleFunc("/PLAY2", PlaySecond).Methods("POST")
	r.HandleFunc("/SAY", Say).Methods("GET")
	r.HandleFunc("/GATHER", Gather).Methods("GET")
	r.HandleFunc("/DTMF", DTMF).Methods("POST")
	r.HandleFunc("/SPEECH", Speech).Methods("GET")
	r.HandleFunc("/speechResponse", SpeechResponse).Methods("POST")
	/*	r.HandleFunc("/DTMF3", DTMF3).Methods("POST")*/
	// Play Say Gather Speech
	r.HandleFunc("/Dial", Dial).Methods("POST")
	r.HandleFunc("/SMS", SMS).Methods("POST")
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
			Language: "en-US",
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

func Gather(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseGather{
		Gather: s.GatherSay{
			NumDigits:   1,
			FinishOnKey: "#",
			Action:      "https://350542707636.ngrok.io/DTMF",
			Say: s.Say{
				Value: "For customer service department, press 1." +
					"For technical support, press 2." +
					"For the sales department, press 3." +
					"Press 4 if you want to return to the main menu.",
				Voice:    "woman",
				Language: "en-US",
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

func Speech(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseGather{
		Gather: s.GatherSay{
			NumDigits:   1,
			FinishOnKey: "#",
			Hints:       "SERVICE",
			Input:       "speech",
			Action:      "https://350542707636.ngrok.io/speechResponse",
			Say: s.Say{
				Value:    "For customer service department, say SERVICE.",
				Voice:    "woman",
				Language: "en-US",
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

func DTMF(w http.ResponseWriter, r *http.Request) {
	println("+++++++++++++++++++++++++++++++++++++")
	w.Header().Set("Content-Type", "application/xml")
	response := "<Response><Say>We are about to connect you with customer service department. Please wait for a moment.</Say></Response>"
	w.Write([]byte(response))
}

func SpeechResponse(w http.ResponseWriter, r *http.Request) {
	println("+++++++++++++++++++++++++++++++++++++")
	w.Header().Set("Content-Type", "application/xml")
	response := "<Response><Say>We are about to connect you with customer service department. Please wait for a moment.</Say></Response>"
	w.Write([]byte(response))
}

func DTMF3(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponseGather{
		Gather: s.GatherSay{
			NumDigits: 1,
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
	inbound := &s.ResponseSms{
		Sms: s.Sms{
			Value: "Hi, the Sms service is working.",
			From:  "(647) 930-8804",
			To:    "(647) 695-6429",
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
