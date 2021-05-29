package config

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Operations int

const (
	MakeCall Operations = iota + 1
	ListCall
	ViewCall
)

type ConfigType struct {
	HttpClient http.Client
	AccountSid string
	AuthToken  string
	From       string
	To         string
	AvayaNumBR string
	AvayaNumCA string
	NumberCA1  string
	NumberBR1  string
	NumberBR2  string
	NumberBR3  string
	ActionUrl  string
	ApiUrl     string
	ApiVersion string
}

func (c ConfigType) GetBaseURL() string {
	return c.ApiUrl + "/" + c.ApiVersion
}

func NewConfig(args []string) (config ConfigType) {
	config = ReadConfig(config)
	config.HttpClient = http.Client{
		Timeout: 60 * time.Second,
	}
	length := len(args)
	if length > 1 {
		println("CLI Arguments")
		config.From = args[0]
		config.To = args[1]
		config.ActionUrl = args[2]
	} else if length == 1 {
		// TODO View Call
	} else {
		println("Default")
		config.From = config.AvayaNumBR
		config.To = config.NumberBR3
	}
	return config
}

func ReadConfig(config ConfigType) ConfigType {
	var configfile = "internal/config/config.ini"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("File configuration "+configfile+" missing: ", configfile)
	}
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
