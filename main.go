package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"gomoko/config"
	"gomoko/utils"
)

func main() {
	httpClient, err := config.NewHttpClient()
	if err != nil {
		panic(err)
	}
	client := twitter.NewClient(httpClient)

	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count:    200,
		TrimUser: twitter.Bool(true),
	})
	if err != nil {
		panic(err)
	}
	//for _, tweet := range tweets {
	//	fmt.Println(tweet.Text)
	//}
	tweetsStrings := utils.TweetToStrings(&tweets)
	var tokens []string
	chain := NewChain()
	for _, tweet := range tweetsStrings {
		tokens = Tokenize(tweet)
		chain.Add(tokens)
	}
	fmt.Println(GenerateTweetText(chain))
}
