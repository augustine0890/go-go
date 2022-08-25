package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *Config
)

type Config struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig(fileName string) (c *Config, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c = new(Config)
	err = json.NewDecoder(file).Decode(c)

	return
}
