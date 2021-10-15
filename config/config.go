package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigValues struct {
	CommandPrefix string `json:"commandPrefix"`
}

type SecretValues struct {
	Token  string       `json:"token"`
	Google GoogleValues `json:"google"`
}

type GoogleValues struct {
	Id string `json:"id"`
	Cx string `json:"cx"`
}

var (
	Config  = new(ConfigValues)
	Secrets = new(SecretValues)
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read config")
	readJSON(wd+"/config/config.json", &Config)
	log.Println("Read secrets")
	readJSON(wd+"/config/secrets.json", &Secrets)
}

func readJSON(fileName string, v interface{}) {
	configFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, v)

	configFile.Close()
}
