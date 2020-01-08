package utils

import (
	"github.com/dghubble/go-twitter/twitter"
	"regexp"
	"strconv"
	"strings"
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

// ２つのstringのsliceを比較し、差分を返す関数
// a := []string{"apple", "banana", "orange"}
// b := []string{"apple", "banana"}
// => ["orange"]
func Difference(a, b []string) []string {
	mapB := make(map[string]struct{}, len(b))
	for _, x := range b {
		mapB[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mapB[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func IntToStringSlice(slice []int64) []string {
	returnSlice := make([]string, len(slice))
	for _, val := range slice {
		returnSlice = append(returnSlice, strconv.Itoa(int(val)))
	}
	return returnSlice
}

func MatchToMyName(text string) bool {
	var matchWords = []string{
		"もこちゃん",
		"もこねえ",
		"もこ姉",
		"もこ姐",
		"もこねぇ",
		"もこあね",
		"もこぴく",
		"曖昧模糊",
	}

	for _, matchWord := range matchWords {
		if strings.Contains(text, matchWord) {
			return true
		}
	}
	return false
}
