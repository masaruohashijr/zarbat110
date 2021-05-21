package services

import (
	"CPaaS/config"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Language int

const (
	Bash Language = iota + 1
	Go
	Node
	Python
)

type AreaCode string

const (
	AC416 AreaCode = "416"
	AC647          = "647"
	AC437          = "437"
	AC506          = "506"
	AC807          = "807"
	AC613          = "613"
	AC702          = "702"
)

type Echo int

const (
	Off Echo = iota
	On
)

type Client struct {
	HttpClient        *http.Client
	BaseURL           *url.URL
	AccountSid        string
	AuthToken         string
	IsAuthenticated   bool
	Authentication    *AuthenticationService
	Messaging         *MessagingService
	Number            *NumberService
	SpeechRecognition *SpeechRecognitionService
	Text2Speech       *Text2SpeechService
	VoiceCall         *VoiceCallService
}

func NewClient(config config.ConfigType) *Client {
	parsedURL, _ := url.Parse(config.ApiUrl)
	c := &Client{
		HttpClient: &config.HttpClient,
		BaseURL:    parsedURL,
		AccountSid: config.AccountSid,
		AuthToken:  config.AuthToken,
	}
	c.BaseURL = parsedURL
	c.Authentication = &AuthenticationService{client: c}
	c.Messaging = &MessagingService{client: c}
	c.Number = &NumberService{client: c}
	c.SpeechRecognition = &SpeechRecognitionService{client: c}
	c.Text2Speech = &Text2SpeechService{client: c}
	return c
}

func (c *Client) NewRequest(method string, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	rel.Path = strings.TrimLeft(rel.Path, "/")
	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.AccountSid, c.AuthToken)
	if err != nil {
		println(err)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	c.HttpClient.Timeout = 60 * time.Second
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
