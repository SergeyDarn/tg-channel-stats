package internal

import (
	"slices"
	"sort"
)

type HashtagMap map[string]Hashtag

func (h HashtagMap) ToSlice() Hashtags {
	var hashtags Hashtags

	for _, hashtag := range h {
		hashtags = append(hashtags, hashtag)
	}

	return hashtags
}

type Hashtags []Hashtag

func (h Hashtags) SortByPopularity(useSingleReaction bool) Hashtags {
	sort.SliceStable(h, func(i, j int) bool {
		iCount := h[i].ReactionsMap.SortCount(useSingleReaction)
		jCount := h[j].ReactionsMap.SortCount(useSingleReaction)

		return iCount > jCount
	})

	return h
}

func (h Hashtags) SortByUnpopularity(useSingleReaction bool) Hashtags {
	sort.SliceStable(h, func(i, j int) bool {
		iCount := h[i].ReactionsMap.SortCount(useSingleReaction)
		jCount := h[j].ReactionsMap.SortCount(useSingleReaction)

		return iCount < jCount
	})

	return h
}

func (h Hashtags) Print() {
	slices.Reverse(h)

	for _, hashtag := range h {
		hashtag.Print()
	}
}

func (h Hashtags) PrintWithReactions() {
	slices.Reverse(h)

	for _, hashtag := range h {
		hashtag.PrintWithReactions()
	}
}
