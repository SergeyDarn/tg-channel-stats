package main

import (
	"flag"

	"github.com/tg-channel-stats/internal"
)

func main() {
	channel := internal.ReadChannelJson()
	posts := channel.GetPosts()

	postsToShow := flag.Int("qty", internal.DefaultPostsToShow, "Number of posts to show. Default is 5")
	getPopular := flag.Bool("popular", true, "Show popular posts")
	getUnpopular := flag.Bool("unpopular", false, "Show unpopular posts")
	flag.Parse()

	if *getUnpopular {
		posts = internal.SortByUnpopularity(posts)
	} else if *getPopular {
		posts = internal.SortByPopularity(posts)
	}

	shortenedPosts := internal.ShortenPosts(posts[0:*postsToShow])
	internal.PrintPosts(shortenedPosts)
}
