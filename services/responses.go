package services

type RestException struct {
	Status  string `xml:"Status"`
	Message string `xml:"Message"`
	Code    string `xml:"Code"`
}

type Response struct {
	IncomingPhoneNumer IncomingPhoneNumer `xml:"IncomingPhoneNumber"`
	RestException      RestException      `xml:"RestException"`
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
	Total                string               `json:"total"`
	PreivousPageUri      string               `json:"previous_page_uri"`
	NumPages             string               `json:"num_pages"`
	IncomingPhoneNumbers []IncomingPhoneNumer `json:"incoming_phone_numbers"`
	Uri                  string               `json:"uri"`
	PageSize             string               `json:"page_size"`
	Start                string               `json:"start"`
	NextPageUri          string               `json:"next_page_uri"`
	LastPageUri          string               `json:"last_page_uri"`
	Page                 string               `json:"page"`
}
