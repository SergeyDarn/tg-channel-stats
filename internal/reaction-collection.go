package internal

import "sort"

type Reactions []Reaction

func (r Reactions) Sort() Reactions {
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Count > r[j].Count
	})

	return r
}

func (r Reactions) ToString() string {
	var reactionsStr string

	for _, reaction := range r {
		reactionsStr += reaction.ToString()
	}

	return reactionsStr
}

type ReactionsMap map[string]Reaction

func (r ReactionsMap) AllCount() int {
	var count int

	for _, reaction := range r {
		count += reaction.Count
	}

	return count
}

func (r ReactionsMap) TopCount() int {
	var topReactionCount int

	for _, reaction := range r {
		if reaction.Count > topReactionCount {
			topReactionCount = reaction.Count
		}
	}

	return topReactionCount
}

func (r ReactionsMap) SortCount(useSingleReaction bool) int {
	if useSingleReaction {
		return r.TopCount()
	}

	return r.AllCount()
}

func (r ReactionsMap) ToSlice() Reactions {
	var reactions Reactions

	for _, reaction := range r {
		reactions = append(reactions, reaction)
	}

	return reactions
}

func (r ReactionsMap) ToStringSorted() string {
	return r.ToSlice().Sort().ToString()
}

func (r ReactionsMap) Add(reactions []Reaction) ReactionsMap {
	for _, reaction := range reactions {
		count := r[reaction.Emoji].Count + reaction.Count

		r[reaction.Emoji] = Reaction{
			Emoji: reaction.Emoji,
			Count: count,
		}
	}

	return r
}
