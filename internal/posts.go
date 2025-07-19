package internal

import (
	"fmt"
	"slices"
	"sort"

	"github.com/charmbracelet/lipgloss"
)

const (
	shortenedMessageLength = 200
	separator              = "-----------------------------------"
	DefaultPostsToShow     = 5

	colorOrange    = lipgloss.Color("#f98b6c")
	colorBlue      = lipgloss.Color("#4a26fd")
	colorPink      = lipgloss.Color("#ef4fa6")
	colorLightBlue = lipgloss.Color("#1e88e5")
	colorTeal      = lipgloss.Color("#03dac6")
)

type Post struct {
	BasePost
	Text string
}

type BasePost struct {
	Id        int
	Type      string
	Date      string
	Reactions []Reaction
}

type Reaction struct {
	Type  string
	Count int
	Emoji string
}

func (p *Post) GetReactionsCount() int {
	var count int

	for _, reaction := range p.Reactions {
		count += reaction.Count
	}

	return count
}

func (p *Post) Shorten() Post {
	text := p.Text
	if len(p.Text) > shortenedMessageLength {
		text = text[0:shortenedMessageLength]
	}

	return Post{
		BasePost: p.BasePost,
		Text:     text,
	}
}

func SortByPopularity(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].GetReactionsCount() > posts[j].GetReactionsCount()
	})

	return posts
}

func SortByUnpopularity(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].GetReactionsCount() < posts[j].GetReactionsCount()
	})

	return posts
}

func ShortenPosts(posts []Post) []Post {
	shortenedPosts := []Post{}

	for _, post := range posts {
		shortenedPosts = append(shortenedPosts, post.Shorten())
	}

	return shortenedPosts
}

func PrintPosts(posts []Post) {
	slices.Reverse(posts)

	for _, post := range posts {
		fmt.Println()
		fmt.Println(separator)
		fmt.Printf(PrepareColorOutput("Id: %d", colorPink), post.Id)
		fmt.Println()
		fmt.Println(post.Text)
		fmt.Printf(
			PrepareColorOutput("Reactions: %d %v", colorOrange),
			post.GetReactionsCount(),
			post.Reactions,
		)
		fmt.Println()
		fmt.Println(separator)
		fmt.Println()
	}
}
