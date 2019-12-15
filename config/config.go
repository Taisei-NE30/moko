package config

import "github.com/BurntSushi/toml"

type Config struct {
	ConsumerToken string `toml:"consumerToken"`
	ConsumerSecret string `toml:"consumerSecret"`
	AccessToken string `toml:"accessToken"`
	AccessSecret string `toml:"accessSecret"`
}

var config Config

func authTwitter() {
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		panic(err)
	}
}
