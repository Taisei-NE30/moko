package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"gomoko/config"
	"gomoko/utils"
	"log"
	"sync"
)

func main() {
	httpClient, err := config.NewHttpClient()
	if err != nil {
		panic(err)
	}
	client := twitter.NewClient(httpClient)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
			Count:    200,
			TrimUser: twitter.Bool(true),
		})
		if err != nil {
			log.Fatal(err)
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
		tweetText := utils.RegexTweet(GenerateTweetText(chain))
		fmt.Println(tweetText)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mentions, _, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{
			Count: 50,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, mention := range mentions {
			fmt.Println(mention.Text)
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

	}()

	wg.Wait()
}
