package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"gomoko/config"
)

func main() {
	httpClient, err := config.NewHttpClient()
	if err != nil {
		panic(err)
	}
	client := twitter.NewClient(httpClient)

	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		panic(err)
	}
	//for _, tweet := range tweets {
	//	fmt.Println(tweet.Text)
	//}

}
