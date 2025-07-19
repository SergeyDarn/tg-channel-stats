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
	flag.Parse()

	if !*getPopular && !*getUnpopular && !*getLongest && !*getShortest {
		fmt.Println("Please use one of the available stat params")
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
