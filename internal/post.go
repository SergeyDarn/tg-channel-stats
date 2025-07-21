package internal

import (
	"fmt"
	"strings"

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
	Id           string
	Text         string
	Hashtags     []string
	ReactionsMap ReactionsMap
}

type BasePost struct {
	Type string
	Date string
}

func (p *Post) HashtagsString() string {
	return strings.Join(p.Hashtags, " ")
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

func (p *Post) ReactionsCount() int {
	return p.ReactionsMap.AllCount()
}

func (p *Post) ReactionsString() string {
	return p.ReactionsMap.ToStringSorted()
}

func (p *Post) ReactionsSortCount(useSingleReaction bool) int {
	return p.ReactionsMap.SortCount(useSingleReaction)
}

func (p *Post) Print() {
	fmt.Println()
	fmt.Println(separator)
	fmt.Printf(PrepareColorOutput("Id: %s", colorPink), p.Id)
	fmt.Println()
	fmt.Println(p.ShortText())
	fmt.Printf(
		PrepareColorOutput("Reactions: %d %v", colorOrange),
		p.ReactionsCount(),
		p.ReactionsString(),
	)
	fmt.Println()
	fmt.Printf(PrepareColorOutput("Word count: %d", colorLightBlue), p.WordCount())
	fmt.Println()
	fmt.Printf(PrepareColorOutput("Hashtags: %s", colorTeal), p.HashtagsString())
	fmt.Println()
	fmt.Println(separator)
	fmt.Println()
}
