package config

import (
	"github.com/BurntSushi/toml"
	"github.com/dghubble/oauth1"
	"net/http"

	//"os"
)

type Config struct {
	ConsumerKey string `toml:"consumerKey"`
	ConsumerSecret string `toml:"consumerSecret"`
	AccessToken string `toml:"accessToken"`
	AccessSecret string `toml:"accessSecret"`
}

var config Config

func NewHttpClient() (*http.Client, error) {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return nil, err
	}
	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	oauthToken := oauth1.NewToken(config.AccessToken, config.AccessSecret)

	/*** production ***/

	//oauthConfig = oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	//oauthToken = oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))

	httpClinet := oauthConfig.Client(oauth1.NoContext, oauthToken)
	return httpClinet, nil
}
