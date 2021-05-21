package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strings"
)

type NumberService struct {
	client *Client
}

func (s *NumberService) ReadBuyNumberResponse(location string) IncomingPhoneNumer {
	xmlFile, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(xmlFile)
	defer xmlFile.Close()
	var response Response
	xml.Unmarshal(byteValue, &response)
	return response.IncomingPhoneNumer
}

func (s *NumberService) BuyNumber(lang Language, areaCode AreaCode, echo Echo) (number IncomingPhoneNumer, err error) {
	apiEndpoint := fmt.Sprintf("/v2/Accounts/%s/IncomingPhoneNumbers", s.client.AccountSid)
	data := url.Values{}
	data.Set("AreaCode", string(areaCode))
	req, err := s.client.NewRequest("POST", apiEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := s.client.Do(req)
	if err != nil {
		println(err.Error())
		var n IncomingPhoneNumer
		return n, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var response Response
	xml.Unmarshal(respBody, &response)
	if echo == On {
		fmt.Println("response Status : ", response.RestException.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(respBody))
	}
	if response.RestException.Status != "400" {
		response.IncomingPhoneNumer.IsBought = true
	}
	return response.IncomingPhoneNumer, nil
}
func (s *NumberService) SetDMTF(number IncomingPhoneNumer, DTMF string, question string) {}

func (s *NumberService) ListNumbers(echo Echo) (NumberAPIResponse, error) {
	apiEndpoint := fmt.Sprintf("v2/Accounts/%s/IncomingPhoneNumbers.json", s.client.AccountSid)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := s.client.Do(req)
	if err != nil {
		println(err.Error())
		n := new(NumberAPIResponse)
		return *n, err
	}
	response := new(NumberAPIResponse)
	json.NewDecoder(resp.Body).Decode(&response)
	if echo == On {
		println("Discovery ====>")
		for _, phone := range response.IncomingPhoneNumbers {
			println(phone.FriendlyName, phone.DateUpdated)
		}
	}
	return *response, nil
}

func PickRandonNumber(numbers []IncomingPhoneNumer) IncomingPhoneNumer {
	return numbers[rand.Intn(len(numbers))]
}
