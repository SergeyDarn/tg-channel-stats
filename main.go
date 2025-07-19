package main

import (
	"flag"

	"github.com/tg-channel-stats/internal"
)

func main() {
	channel := internal.ReadChannelJson()
	posts := channel.GetPosts()

	postsToShow := flag.Int("qty", internal.DefaultPostsToShow, "")
	getPopular := flag.Bool("popular", false, "")
	getUnpopular := flag.Bool("unpopular", false, "")
	flag.Parse()

	if *getPopular {
		posts = internal.SortByPopularity(posts)
	} else if *getUnpopular {
		posts = internal.SortByUnpopularity(posts)
	}

	shortenedPosts := internal.ShortenPosts(posts[0:*postsToShow])
	internal.PrintPosts(shortenedPosts)
}
