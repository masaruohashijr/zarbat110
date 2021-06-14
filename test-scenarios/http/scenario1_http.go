package main

import (
	calls "CPaaS/cmd/calls/handlers"
	"CPaaS/cmd/gather/handlers"
	say "CPaaS/cmd/say/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/MakeCall", calls.MakeCallHandler).Methods("POST")
	r.HandleFunc("/SaySomething", say.SaySomethingHandler).Methods("POST")
	r.HandleFunc("/Gather", handlers.GatherhHandler).Methods("POST")
	r.HandleFunc("/SpeechResult", handlers.SpeechResultHandler).Methods("POST", "GET")
	//r.HandleFunc("/Ola", OlaMundo).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}

func OlaMundo(w http.ResponseWriter, r *http.Request) {
	println("*************************************")
	println("OLÁ MUNDO")
	println("*************************************")
}
