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

type PostsInfo struct {
	PostCount int
	WordCount int
}

func (p *Post) ReactionCount() int {
	var count int

	for _, reaction := range p.Reactions {
		count += reaction.Count
	}

	return count
}

func (p *Post) ShortText() string {
	if len(p.Text) <= shortenedMessageLength {
		return p.Text
	}

	return p.Text[0:shortenedMessageLength]
}

func (p *Post) WordCount() int {
	return CountWords(p.Text)
}

func (p *Post) Print() {
	fmt.Println()
	fmt.Println(separator)
	fmt.Printf(PrepareColorOutput("Id: %d", colorPink), p.Id)
	fmt.Println()
	fmt.Println(p.ShortText())
	fmt.Printf(
		PrepareColorOutput("Reactions: %d %v", colorOrange),
		p.ReactionCount(),
		p.Reactions,
	)
	fmt.Println()
	fmt.Printf(PrepareColorOutput("Word count: %d", colorLightBlue), p.WordCount())
	fmt.Println()
	fmt.Println(separator)
	fmt.Println()
}

func SortByPopularity(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].ReactionCount() > posts[j].ReactionCount()
	})

	return posts
}

func SortByUnpopularity(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].ReactionCount() < posts[j].ReactionCount()
	})

	return posts
}

func SortLongest(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].WordCount() > posts[j].WordCount()
	})

	return posts
}

func SortShortest(posts []Post) []Post {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].WordCount() < posts[j].WordCount()
	})

	return posts
}

func GetAverageWordCount(posts []Post, minWordCount int) int {
	info := GetPostsInfo(posts, minWordCount)

	return int(info.WordCount / info.PostCount)
}

func GetPostCount(posts []Post, minWordCount int) int {
	info := GetPostsInfo(posts, minWordCount)

	return info.PostCount
}

func GetPostsInfo(posts []Post, minWordCount int) PostsInfo {
	info := PostsInfo{}

	for _, post := range posts {
		postWordCount := post.WordCount()
		if postWordCount < minWordCount {
			continue
		}

		info.WordCount += postWordCount
		info.PostCount++
	}

	return info
}

func PrintPosts(posts []Post) {
	slices.Reverse(posts)

	for _, post := range posts {
		post.Print()
	}
}
