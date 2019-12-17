package utils

import "github.com/dghubble/go-twitter/twitter"

func TweetToStrings(tweets *[]twitter.Tweet) []string {
	returnStrings := make([]string, 0)
	for _, tweet := range *tweets {
		returnStrings = append(returnStrings, tweet.Text)
	}
	return returnStrings
}
