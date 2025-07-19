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

func (p *Post) ShortText() string {
	text := p.Text
	if len(p.Text) > shortenedMessageLength {
		text = text[0:shortenedMessageLength]
	}
	return text
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

func GetAverageWordCount(posts []Post, ignoreBelowCount int) int {
	var wordCount int
	var postsCount int

	for _, post := range posts {
		postWordCount := CountWords(post.Text)
		if postWordCount < ignoreBelowCount {
			continue
		}

		wordCount += CountWords(post.Text)
		postsCount++
	}

	return int(wordCount / postsCount)
}

func SortLongest(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return CountWords(posts[i].Text) > CountWords(posts[j].Text)
	})

	return posts
}

func SortShortest(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return CountWords(posts[i].Text) < CountWords(posts[j].Text)
	})

	return posts
}

func PrintPosts(posts []Post) {
	slices.Reverse(posts)

	for _, post := range posts {
		fmt.Println()
		fmt.Println(separator)
		fmt.Printf(PrepareColorOutput("Id: %d", colorPink), post.Id)
		fmt.Println()
		fmt.Println(post.ShortText())
		fmt.Printf(
			PrepareColorOutput("Reactions: %d %v", colorOrange),
			post.GetReactionsCount(),
			post.Reactions,
		)
		fmt.Println()
		fmt.Printf(PrepareColorOutput("Word count: %d", colorLightBlue), CountWords(post.Text))
		fmt.Println()
		fmt.Println(separator)
		fmt.Println()
	}
}
