package services

import "encoding/xml"

type RestException struct {
	Status  string `xml:"Status"`
	Message string `xml:"Message"`
	Code    string `xml:"Code"`
}

type IncomingPhoneNumer struct {
	MmsUrl               string       `xml:"MmsUrl" json:"mms_url"`
	DateUpdated          string       `xml:"DateUpdated" json:"date_updated"`
	MmsMethod            string       `xml:"MmsMethod" json:"mms_method"`
	VoiceUrl             string       `xml:"VoiceUrl" json:"voice_url"`
	VoiceFallbackMethod  string       `xml:"VoiceFallbackMethod" json:"voice_fallback_method"`
	MmsFallbackMethod    string       `xml:"MmsFallbackMethod" json:"mms_fallback_method"`
	Sid                  string       `xml:"Sid" json:"sid"`
	HeartbeatMethod      string       `xml:"HeartbeatMethod" json:"heartbeat_method"`
	MmsFallbackUrl       string       `xml:"MmsFallbackUrl" json:"mms_fallback_url"`
	StatusCallbackMethod string       `xml:"StatusCallbackMethod" json:"status_callback_method"`
	VoiceFallbackUrl     string       `xml:"VoiceFallbackUrl" json:"voice_fallback_url"`
	Capabilities         Capabilities `xml:"Capabilities" json:"capabilities"`
	PhoneNumber          string       `xml:"PhoneNumber" json:"phone_number"`
	HangupCallback       string       `xml:"HangupCallback" json:"hangup_callback"`
	HangupCallbackMethod string       `xml:"HangupCallbackMethod" json:"hangup_callback_method"`
	HeartbeatUrl         string       `xml:"HeartbeatUrl" json:"heartbeat_url"`
	SmsUrl               string       `xml:"SmsUrl" json:"sms_url"`
	VoiceMethod          string       `xml:"VoiceMethod" json:"voice_method"`
	Type                 string       `xml:"Type" json:"type"`
	VoiceCallerIdLookup  string       `xml:"VoiceCallerIdLookup" json:"voice_caller_id_lookup"`
	FriendlyName         string       `xml:"FriendlyName" json:"friendly_name"`
	Uri                  string       `xml:"Uri" json:"uri"`
	SmsFallbackUrl       string       `xml:"SmsFallbackUrl" json:"sms_fallback_url"`
	AccountSid           string       `xml:"AccountSid" json:"account_sid"`
	SmsMethod            string       `xml:"SmsMethod" json:"sms_method"`
	ApiVersion           string       `xml:"ApiVersion" json:"api_version"`
	SmsFallbackMethod    string       `xml:"SmsFallbackMethod" json:"sms_fallback_method"`
	NextRenewalDate      string       `xml:"NextRenewalDate" json:"next_renewal_date"`
	DateCreated          string       `xml:"DateCreated" json:"date_created"`
	StatusCallback       string       `xml:"StatusCallback" json:"status_callback"`
	IsBought             bool
}

type Capabilities struct {
	Voice string `xml:"Voice" json:"voice"`
	Sms   string `xml:"Sms" json:"sms"`
	Mms   string `xml:"Mms" json:"mms"`
}

type NumberAPIResponse struct {
	FirstPageUri         string               `json:"first_page_uri"`
	End                  string               `json:"end"`
	Total                int                  `json:"total"`
	PreivousPageUri      string               `json:"previous_page_uri"`
	NumPages             int                  `json:"num_pages"`
	IncomingPhoneNumbers []IncomingPhoneNumer `json:"incoming_phone_numbers"`
	Uri                  string               `json:"uri"`
	PageSize             int                  `json:"page_size"`
	Start                string               `json:"start"`
	NextPageUri          string               `json:"next_page_uri"`
	LastPageUri          string               `json:"last_page_uri"`
	Page                 string               `json:"page"`
}

type Say struct {
	Value    string `xml:",chardata"`
	Voice    string `xml:"voice,attr"`
	Language string `xml:"language,attr"`
	Loop     int    `xml:"loop,attr"`
}

type Sms struct {
	Value                string `xml:",chardata"`
	To                   string `xml:"to"`
	From                 string `xml:"from"`
	Action               string `xml:"action"`
	Method               string `xml:"method"`
	StatusCallback       string `xml:"statusCallback"`
	StatusCallbackMethod string `xml:"statusCallbackMethod"`
}

type Mms struct {
	Value          string `xml:",chardata"`
	To             string `xml:"to"`
	From           string `xml:"from"`
	Action         string `xml:"action"`
	Method         string `xml:"method"`
	StatusCallback string `xml:"statusCallback"`
	MediaUrl       string `xml:"mediaUrl"`
}

type Dial struct {
	Value             string `xml:",chardata"`
	Action            string `xml:"action"`
	Method            string `xml:"method"`
	Timeout           int    `xml:"timeout"`
	HangupOnStar      string `xml:"hangupOnStar"`
	TimeLimit         int    `xml:"timeLimit"`
	CallerId          string `xml:"callerId"`
	HideCallerId      string `xml:"hideCallerId"`
	CallerName        string `xml:"callerName"`
	DialMusic         string `xml:"dialMusic"`
	CallbackUrl       string `xml:"callbackUrl"`
	CallbackMethod    string `xml:"callbackMethod"`
	ConfirmSound      string `xml:"confirmSound"`
	DigitsMatch       string `xml:"digitsMatch"`
	StraightToVm      string `xml:"straightToVm"`
	HeartbeatUrl      string `xml:"heartbeatUrl"`
	HeartbeatMethod   string `xml:"heartbeatMethod"`
	ForwardedFrom     string `xml:"forwardedFrom"`
	Record            string `xml:"record"`
	RecordDirection   string `xml:"recordDirection"`
	RecordCallbackUrl string `xml:"recordCallbackUrl"`
	RecordLifetime    string `xml:"recordLifetime"`
	RecordFormat      string `xml:"recordFormat"`
	IfMachine         string `xml:"ifMachine"`
	IfMachineUrl      string `xml:"ifMachineUrl"`
	IfMachineMethod   string `xml:"ifMachineMethod"`
	OutboundAction    string `xml:"outboundAction"`
}

type Play struct {
	Value  string `xml:",chardata"`
	Loop   int    `xml:"loop,attr"`
	Method int    `xml:"method,attr"`
}

type Pause struct {
	Length int `xml:"length,attr"`
}

type Redirect struct {
	Method string `xml:"method,attr"`
}

type Ping struct {
	Value  string `xml:",chardata"`
	Method string `xml:"method,attr"`
}

type Reject struct {
	Reason string `xml:"reason,attr"`
}

type GetSpeech struct {
	Say               Say    `xml:"Say"`
	Play              string `xml:"Play"`
	PlayLastRecording string `xml:"PlayLastRecording"`
	Pause             Pause  `xml:"Pause"`
	Action            string `xml:"action,attr"`
	Method            string `xml:"method,attr"`
	Timeout           int    `xml:"timeout,attr"`
	PlayBeep          string `xml:"playBeep,attr"`
	Grammar           int    `xml:"grammar,attr"`
}

type Gather struct {
	Say               Say    `xml:"Say"`
	Play              string `xml:"Play"`
	PlayLastRecording string `xml:"PlayLastRecording"`
	Pause             Pause  `xml:"Pause"`
	Input             string `xml:"input,attr"`
	Hints             string `xml:"hints,attr"`
	Language          string `xml:"language,attr"`
	Action            string `xml:"action,attr"`
	Method            string `xml:"method,attr"`
	Timeout           int    `xml:"timeout,attr"`
	FinishOnKey       string `finishOnKey:"input,attr"`
	NumDigits         int    `xml:"numDigits,attr"`
}

type Hangup struct {
	Schedule string `xml:"schedule,attr"`
	Reason   string `xml:"reason,attr"`
}

type Number struct {
	SendDigits      string `xml:"sendDigits,attr"`
	Method          string `xml:"method,attr"`
	Url             int    `xml:"url,attr"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr"`
}

type User struct {
	SendDigits string `xml:"sendDigits,attr"`
	Params     string `xml:"params,attr"`
}

type Response struct {
	XMLName            xml.Name           `xml:"Response"`
	Say                Say                `xml:"Say"`
	Play               Play               `xml:"Play"`
	Answer             string             `xml:"Answer"`
	PlayLastRecording  string             `xml:"PlayLastRecording"`
	Gather             Gather             `xml:"Gather"`
	Dial               Dial               `xml:"Dial"`
	Hangup             Hangup             `xml:"Hangup"`
	Redirect           Redirect           `xml:"Redirect"`
	Ping               Ping               `xml:"Ping"`
	Reject             Reject             `xml:"Reject"`
	Pause              Pause              `xml:"Pause"`
	Sms                Sms                `xml:"Sms"`
	Mms                Mms                `xml:"Mms"`
	IncomingPhoneNumer IncomingPhoneNumer `xml:"IncomingPhoneNumber"`
	RestException      RestException      `xml:"RestException"`
}
