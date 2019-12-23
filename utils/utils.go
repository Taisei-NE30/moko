package utils

import (
	"github.com/dghubble/go-twitter/twitter"
	"regexp"
)

func TweetToStrings(tweets *[]twitter.Tweet) []string {
	returnStrings := make([]string, 0)
	for _, tweet := range *tweets {
		if tweet.User.Protected {
			continue
		}
		returnStrings = append(returnStrings, tweet.Text)
	}
	return returnStrings
}

func RegexTweet(text string) string {
	text = removeURL(text)
	text = removeRT(text)
	text = removeReply(text)
	return text
}

func removeURL(text string) string {
	re := regexp.MustCompile(`(http|https)://t.co/\w+`)
	removedText := re.ReplaceAllString(text, "")
	return removedText
}

func removeRT(text string) string {
	re := regexp.MustCompile(`^RT\s@\w+:\s`)
	removedText := re.ReplaceAllString(text, "")
	return removedText
}

func removeReply(text string) string {
	re := regexp.MustCompile(`@\w+`)
	removedText := re.ReplaceAllString(text, "")
	return removedText
}
