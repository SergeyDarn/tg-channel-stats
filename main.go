package main

import (
	"flag"
	"fmt"

	"github.com/tg-channel-stats/internal"
)

func main() {
	channel := internal.ReadChannelJson()
	posts := channel.GetPosts()

	postsToShow := flag.Int("qty", internal.DefaultPostsToShow, "Number of posts to show. Default is 5")
	getPopular := flag.Bool("popular", false, "Show popular posts")
	getUnpopular := flag.Bool("unpopular", false, "Show unpopular posts")
	getLongest := flag.Bool("longest", false, "Show longest posts")
	getShortest := flag.Bool("shortest", false, "Show shortes posts")
	getPostCount := flag.Bool("post-count", false, "How many posts you posted")
	getAverageWordCount := flag.Bool("average-word-count", false, "Show average post word count")
	minWordCount := flag.Int("min-word-count", 0, "Count only posts above certain word count")
	flag.Parse()

	flagPassed := *getPopular || *getUnpopular || *getLongest || *getShortest ||
		*getPostCount || *getAverageWordCount

	if !flagPassed {
		fmt.Println("Please use one of the available stat params")
		return
	}

	if *getPostCount {
		postCount := internal.GetPostCount(posts, *minWordCount)
		fmt.Println(postCount)
		return
	}

	if *getAverageWordCount {
		wordCount := internal.GetAverageWordCount(posts, *minWordCount)
		fmt.Println(wordCount)
		return
	}

	if *getPopular {
		posts = internal.SortByPopularity(posts)
	} else if *getUnpopular {
		posts = internal.SortByUnpopularity(posts)
	}

	if *getLongest {
		posts = internal.SortLongest(posts)
	} else if *getShortest {
		posts = internal.SortShortest(posts)
	}

	internal.PrintPosts(posts[0:*postsToShow])
}
