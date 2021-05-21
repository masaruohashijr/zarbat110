package services

type Text2SpeechService struct {
	client *Client
}

func (s *Text2SpeechService) Say(number IncomingPhoneNumer, mp3Sound string) {}
