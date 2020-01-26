package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/mb-14/gomarkov"
	"gomoko/config"
	"gomoko/utils"
	"html"
	"log"
	"strconv"
	"sync"
)

func main() {
	httpClient, err := config.NewHttpClient()
	if err != nil {
		panic(err)
	}
	client := twitter.NewClient(httpClient)
	wg := &sync.WaitGroup{}
	chainCh := make(chan *gomarkov.Chain)
	twCh := make(chan []twitter.Tweet)

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

		twCh <- tweets
		//for _, tweet := range tweets {
		//	fmt.Println(tweet.Text)
		//}
		tweetsStrings := utils.TweetToStrings(&tweets)

		var tokens []string
		chain := NewChain()

		for _, tweet := range tweetsStrings {
			tweet = html.UnescapeString(tweet)
			tweet = utils.RegexTweet(tweet)
			tokens = Tokenize(tweet)
			chain.Add(tokens)
		}
		chainCh <- chain

		//fmt.Println(GenerateTweetText(chain))
		_, _, err = client.Statuses.Update(GenerateTweetText(chain), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// メンションに返信する
	wg.Add(1)
	go func() {
		defer wg.Done()

		myTweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName: "mokonee0607",
			Count:      100,
		})
		if err != nil {
			log.Fatal(err)
		}

		var replyedIds []int64
		for _, myTweet := range myTweets {
			if myTweet.InReplyToStatusID != 0 {
				replyedIds = append(replyedIds, myTweet.InReplyToStatusID)
			}
		}

		mentions, _, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{
			Count: 20,
		})
		if err != nil {
			log.Fatal(err)
		}

		chain := <-chainCh

		shouldReply := true
		for _, mention := range mentions {
			for _, replyedId := range replyedIds {
				if mention.ID == replyedId {
					shouldReply = false
				}
			}
			if shouldReply {
				replyUser := mention.User.ScreenName
				tweetText := utils.RegexTweet(GenerateTweetText(chain))
				//fmt.Printf("@%s %s\n", replyUser, tweetText)
				_, _, err := client.Statuses.Update(fmt.Sprintf("@%s %s", replyUser, tweetText), &twitter.StatusUpdateParams{
					InReplyToStatusID: mention.ID,
				})
				if err != nil {
					log.Fatal(err)
				}
			}
			shouldReply = true
		}

	}()

	// 自動フォロバ
	wg.Add(1)
	go func() {
		defer wg.Done()

		followerIDs, _, err := client.Followers.IDs(nil)
		if err != nil {
			log.Fatal(err)
		}
		stringFollowerIDs := utils.IntToStringSlice(followerIDs.IDs)

		friendIDs, _, err := client.Friends.IDs(nil)
		if err != nil {
			log.Fatal(err)
		}
		stringFriendIDs := utils.IntToStringSlice(friendIDs.IDs)

		ffDiffIDs := utils.Difference(stringFollowerIDs, stringFriendIDs)

		for _, id := range ffDiffIDs {
			intID, _ := strconv.Atoi(id)
			_, _, err := client.Friendships.Create(&twitter.FriendshipCreateParams{
				UserID: int64(intID),
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	// 自動いいね
	wg.Add(1)
	go func() {
		defer wg.Done()
		tweets := <-twCh

		for _, tweet := range tweets {
			if utils.MatchToMyName(tweet.Text) {
				_, _, err := client.Favorites.Create(&twitter.FavoriteCreateParams{
					ID: tweet.ID,
				})
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	wg.Wait()
}
