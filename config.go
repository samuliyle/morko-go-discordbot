package main

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
	Token string `json:"token"`
}

var (
	config  = new(ConfigValues)
	secrets = new(SecretValues)
)

func init() {
	log.Println("Read config")
	readJSON("config.json", &config)
	log.Println("Read secrets")
	readJSON("secrets.json", &secrets)
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
