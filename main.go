package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ConsumerToken string `toml:"consumerToken"`
	ConsumerSecret string `toml:"consumerSecret"`
	AccessToken string `toml:"accessToken"`
	AccessSecret string `toml:"accessSecret"`
}

func main() {
	var config Config
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		panic(err)
	}
}

