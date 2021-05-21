package main

import (
	c "CPaaS/config"
	s "CPaaS/services"
	"fmt"
	"strings"
)

func main() {
	config := c.NewConfig()
	cpaas := s.NewClient(config)
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
	}
}
