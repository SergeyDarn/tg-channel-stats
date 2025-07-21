package internal

import (
	"fmt"
	"strings"
)

type Hashtag struct {
	Text         string
	PostsIds     []string
	ReactionsMap ReactionsMap
}

func (h *Hashtag) PostsCount() int {
	return len(h.PostsIds)
}

func (h *Hashtag) PostsIdsString() string {
	return strings.Join(h.PostsIds, ", ")
}

func (h *Hashtag) ReactionsCount() int {
	return h.ReactionsMap.AllCount()
}

func (h *Hashtag) TopReactionCount() int {
	return h.ReactionsMap.TopCount()
}

func (h *Hashtag) SortReactionCount(useSingleReaction bool) int {
	return h.ReactionsMap.SortCount(useSingleReaction)
}

func (h *Hashtag) Print() {
	fmt.Println()
	fmt.Println(PrepareColorOutput(h.Text, colorTeal))
	fmt.Printf(
		"Posts (%d): %s",
		h.PostsCount(),
		h.PostsIdsString(),
	)
	fmt.Println()
}

func (h *Hashtag) PrintWithReactions() {
	h.Print()
	fmt.Printf(
		PrepareColorOutput("Reactions (%d): %s", colorOrange),
		h.ReactionsMap.AllCount(),
		h.ReactionsMap.ToStringSorted(),
	)
	fmt.Println()
}
