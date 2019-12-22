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
	ch := make(chan []twitter.Tweet)

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
		ch <- tweets
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
	// メンションに返信する
	go func() {
		defer wg.Done()

		myTweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName: "mokonee0607",
			Count:      10,
		})
		if err != nil {
			log.Fatal(err)
		}

		var replyedIds []int64
		for _, myTweet := range myTweets {
			if myTweet.InReplyToStatusID != 0 {
				replyedIds = append(replyedId, myTweet.InReplyToStatusID)
			}
		}

		mentions, _, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{
			Count: 20,
		})
		if err != nil {
			log.Fatal(err)
		}

		tweets := <-ch

		for _, mention := range mentions {
			for _, replyedId := range replyedIds {
				if mention.ID != replyedId {

				}
			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

	}()

	wg.Wait()
}
