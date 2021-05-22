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
	cpaas.Authentication.Authenticate(s.Off)
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

	r.HandleFunc("/play1", PlayFirst).Methods("POST")
	r.HandleFunc("/play2", PlaySecond).Methods("POST")
	r.HandleFunc("/say", Say).Methods("POST")
	r.HandleFunc("/answer", Answer).Methods("POST")
	r.HandleFunc("/DTMF1", DTMF1).Methods("POST")
	r.HandleFunc("/DTMF2", DTMF2).Methods("POST")
	r.HandleFunc("/DTMF3", DTMF3).Methods("POST")
	r.HandleFunc("/Dial", Dial).Methods("POST")
	r.HandleFunc("/SMS", SMS).Methods("POST")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err)
		http.ListenAndServe(addr, nil)
	}
}

func PlayFirst(w http.ResponseWriter, r *http.Request) {
	inbound := &s.ResponsePlay{
		Play: "https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/Coldplay/Paradise.mp3",
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
	inbound := &s.Response{
		Play: s.Play{
			Value: "https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/Alok/Hear_Me_Now.mp3",
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
	inbound := &s.Response{
		Say: s.Say{
			Value:    "How are you?",
			Voice:    "woman",
			Language: "en-us",
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

func Answer(w http.ResponseWriter, r *http.Request) {
	inbound := &s.Response{
		Gather: s.Gather{
			Method:   "POST",
			Input:    "speech",
			Hints:    "How are you?",
			Language: "es-US",
			Say: s.Say{
				Value:    "I am good",
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

func DTMF1(w http.ResponseWriter, r *http.Request) {
	inbound := &s.Response{
		Gather: s.Gather{
			Method:      "POST",
			NumDigits:   1,
			FinishOnKey: "1",
			Input:       "dtmf",
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
	inbound := &s.Response{
		Gather: s.Gather{
			Method:      "POST",
			NumDigits:   1,
			FinishOnKey: "2",
			Input:       "dtmf",
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
	inbound := &s.Response{
		Gather: s.Gather{
			Method:      "POST",
			NumDigits:   1,
			FinishOnKey: "3",
			Input:       "dtmf",
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

func PlayPauseRedirec(w http.ResponseWriter, r *http.Request) {
	//2. Using Play, Pause, Redirect to play a file then load new inbound XML to play another file
	inbound := &s.Response{
		Gather: s.Gather{
			Method: "POST",
			Play:   "https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/Alok/Hear_Me_Now.mp3",
			Pause: s.Pause{
				Length: 2,
			},
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
