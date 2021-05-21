package main

import (
	c "CPaaS/config"
	s "CPaaS/services"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	config := c.NewConfig()
	cpaas := s.NewClient(config)
	// echo := s.On
	cpaas.Authentication.Authenticate(s.Off)
	if cpaas.Authentication.IsAuthenticated() {
		boughtNumber := cpaas.Number.ReadBuyNumberResponse(config.ResultBuyNumber)
		fmt.Printf("The number %s was last updated in %s.\n", strings.TrimSpace(boughtNumber.FriendlyName), strings.TrimSpace(boughtNumber.DateUpdated))
		phoneNumber, _ := cpaas.Number.BuyNumber(s.Bash, s.AC647, s.Off)
		if !phoneNumber.IsBought {
			var numbers []s.IncomingPhoneNumer
			numberAPIResponse, _ := cpaas.Number.ListNumbers(s.On)
			numbers = numberAPIResponse.IncomingPhoneNumbers
			phoneNumber = s.PickRandonNumber(numbers)
		}
		fmt.Printf("The number %s was purchased in %s.\n", strings.TrimSpace(phoneNumber.FriendlyName), strings.TrimSpace(phoneNumber.DateUpdated))
		cpaas.VoiceCall.SetReceivingCalls(phoneNumber, config.MP3)
		cpaas.Text2Speech.Say(phoneNumber, "Hello World")
		cpaas.Number.SetDMTF(phoneNumber, config.DTMF1, "Good morning")
		cpaas.Number.SetDMTF(phoneNumber, config.DTMF2, "How are you")
		cpaas.Number.SetDMTF(phoneNumber, config.DTMF3, "How do yo do")
		cpaas.SpeechRecognition.SetRecognization("How are you", "I am good")
		cpaas.VoiceCall.PlaceVoiceCall(config.ToNumber)
		fmt.Printf("The number %s was purchased in %s.\n", strings.TrimSpace(phoneNumber.FriendlyName), strings.TrimSpace(phoneNumber.DateUpdated))
	}
	r := mux.NewRouter()
	http.Handle("/", r)
	addr := ":5000"

	r.HandleFunc("/play", PlayFirst).Methods("POST")
	r.HandleFunc("/say", Say).Methods("POST")
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
	inbound := &s.InboundXMLResponse{
		&Response{
			"Play":"https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/U2/Beautiful_Day.mp3"
		}
	}
	x, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(inbound)
	return
}

func PlaySecond(w http.ResponseWriter, r *http.Request) {
	inbound := &s.InboundXMLResponse{
		&Response{
			"Play":"https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/Alok/Hear_Me_Now.mp3"
		}
	}
	x, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(inbound)
	return
}
func Dial(w http.ResponseWriter, r *http.Request) {
	inbound := &s.InboundXMLResponse{
		&Response{
			"Dial":"https://teresadapraiamidis.com/Mp3/Musicas/Rock_Internacional/U2/Beautiful_Day.mp3"
		}
	}
	x, err := xml.MarshalIndent(inbound, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(inbound)
	return
}
func Say(w http.ResponseWriter, r *http.Request) {
	return
}
func DTMF1(w http.ResponseWriter, r *http.Request) {
	return
}
func DTMF2(w http.ResponseWriter, r *http.Request) {
	return
}
func DTMF3(w http.ResponseWriter, r *http.Request) {
	return
}
func SMS(w http.ResponseWriter, r *http.Request) {
	return
}
