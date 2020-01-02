package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/mb-14/gomarkov"
	"gomoko/config"
	"gomoko/utils"
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
			Count:      10,
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

		for _, mention := range mentions {
			for _, replyedId := range replyedIds {
				if mention.ID != replyedId {
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
			}
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

	}()

	wg.Wait()
}
