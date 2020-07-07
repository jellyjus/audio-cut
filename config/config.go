package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configsPath = "./"

type config struct {
	StaticPath string `json:"static_path"`
	Vk         vk     `json:"vk"`
}

type vk struct {
	AppID       string `json:"app_id"`
	SecretKey   string `json:"secret_key"`
	ServiceKey  string `json:"service_key"`
	AccessToken string `json:"access_token"`
	RedirectURI string `json:"redirect_uri"`
}

var Config = config{}

func ParseConfig(env string) error {
	var path string
	config := os.Getenv("CONFIG")
	if config == "" {
		path = fmt.Sprintf("%sconfig.%s.json", configsPath, env)
	} else {
		path = config
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}
	return nil
}
