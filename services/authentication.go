package services

import (
	"fmt"
	"io/ioutil"
)

type AuthenticationService struct {
	client     *Client
	accountSid string
	authToken  string
}

func (s *AuthenticationService) SetBasicAuth(accountSid string, authToken string) {
	s.accountSid = accountSid
	s.authToken = authToken
}

func (s *AuthenticationService) SetAuthenticated() {
	return
}

func (s *AuthenticationService) IsAuthenticated() bool {
	return s.client.IsAuthenticated
}

func (s *AuthenticationService) Authenticate(echo Echo) (bool, error) {
	apiEndpoint := fmt.Sprintf("v2/Accounts/%s.json", s.client.AccountSid)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := s.client.Do(req)
	if err != nil {
		println(err.Error())
		return false, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	if echo == On {
		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(respBody))
	}
	s.client.IsAuthenticated = true
	return true, nil
}
