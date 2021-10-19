package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigValues struct {
	CommandPrefix string         `json:"commandPrefix"`
	Database      databaseValues `json:"database"`
}

type SecretValues struct {
	Token    string          `json:"token"`
	Google   googleSecrets   `json:"google"`
	Database databaseSecrets `json:"database"`
}

type databaseValues struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type googleSecrets struct {
	Id string `json:"id"`
	Cx string `json:"cx"`
}

type databaseSecrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValue, v)

}
