package config

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type ConfigType struct {
	HttpClient      http.Client
	ApiUrl          string
	AccountSid      string
	AuthToken       string
	MP3             string
	DTMF1           string
	DTMF2           string
	DTMF3           string
	ToNumber        string
	ResultBuyNumber string
}

func NewConfig() (config ConfigType) {
	config = ReadConfig(config)
	config.HttpClient = http.Client{
		Timeout: 60 * time.Second,
	}
	return config
}

func ReadConfig(config ConfigType) ConfigType {
	var configfile = "config/config.ini"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("File configuration "+configfile+" missing: ", configfile)
	}
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
