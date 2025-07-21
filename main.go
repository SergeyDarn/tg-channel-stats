package main

import (
	"flag"
	"fmt"

	"github.com/tg-channel-stats/internal"
)

func main() {
	qty := flag.Int("qty", internal.DefaultPostsToShow, "Number of posts or hashtags to show. Default is 5")
	getPopularPosts := flag.Bool("popular", false, "Show popular posts")
	getUnpopularPosts := flag.Bool("unpopular", false, "Show unpopular posts")
	useSingleReaction := flag.Bool("single-reaction", false, "Calculate posts' popularity by a single top reaction (default is total reactions' count)")
	getLongestPosts := flag.Bool("longest", false, "Show longest posts")
	getShortestPosts := flag.Bool("shortest", false, "Show shortest posts")

	getHashtags := flag.Bool("hashtags", false, "Show all hashtags")
	getPopularHashtags := flag.Bool("hashtags-popular", false, "Show all hashtags, sorted by popularity")
	getUnpopularHashtags := flag.Bool("hashtags-unpopular", false, "Show all hashtags, sorted by unpopularity")

	getPostCount := flag.Bool("post-count", false, "How many posts you posted")
	getAverageWordCount := flag.Bool("average-word-count", false, "Show average word count per post")
	minWordCount := flag.Int("min-word-count", 0, "Count only posts above certain word count")
	flag.Parse()

	flagPassed := *getPopularPosts || *getUnpopularPosts || *getLongestPosts || *getShortestPosts ||
		*getPostCount || *getAverageWordCount || *getHashtags || *getPopularHashtags || *getUnpopularHashtags

	if !flagPassed {
		fmt.Println("Please use one of the available stat params")
		return
	}

	channel := internal.ReadChannelJson()
	posts, hashtagMap := channel.ProcessedData()
	hashtags := hashtagMap.ToSlice()

	if *getHashtags {
		hashtags.Print()
		return
	}

	if *getPopularHashtags {
		sortedHashtags := hashtags.SortByPopularity(*useSingleReaction)
		sortedHashtags[0:*qty].PrintWithReactions()
		return
	}

	if *getUnpopularHashtags {
		sortedHashtags := hashtags.SortByUnpopularity(*useSingleReaction)
		sortedHashtags[0:*qty].PrintWithReactions()
		return
	}

	if *getPostCount {
		fmt.Println(
			posts.Count(*minWordCount),
		)
		return
	}

	if *getAverageWordCount {
		fmt.Println(
			posts.AverageWordCount(*minWordCount),
		)
		return
	}

	if *getPopularPosts {
		posts = posts.SortByPopularity(*useSingleReaction)
	} else if *getUnpopularPosts {
		posts = posts.SortByUnpopularity(*useSingleReaction)
	}

	if *getLongestPosts {
		posts = posts.SortLongest()
	} else if *getShortestPosts {
		posts = posts.SortShortest()
	}

	posts[0:*qty].Print()
}
