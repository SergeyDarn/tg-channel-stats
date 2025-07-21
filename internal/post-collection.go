package internal

import (
	"slices"
	"sort"
)

type Posts []Post

type PostsCounts struct {
	PostCount int
	WordCount int
}

func (p Posts) SortByPopularity(useSingleReaction bool) Posts {
	sort.SliceStable(p, func(i, j int) bool {
		postICount := p[i].ReactionsSortCount(useSingleReaction)
		postJCount := p[j].ReactionsSortCount(useSingleReaction)

		return postICount > postJCount
	})

	return p
}

func (p Posts) SortByUnpopularity(useSingleReaction bool) Posts {
	sort.SliceStable(p, func(i, j int) bool {
		postICount := p[i].ReactionsSortCount(useSingleReaction)
		postJCount := p[j].ReactionsSortCount(useSingleReaction)

		return postICount < postJCount
	})

	return p
}

func (p Posts) SortLongest() Posts {
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].WordCount() > p[j].WordCount()
	})

	return p
}

func (p Posts) SortShortest() Posts {
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].WordCount() < p[j].WordCount()
	})

	return p
}

func (p Posts) AverageWordCount(minWordCount int) int {
	counts := p.Counts(minWordCount)

	return int(counts.WordCount / counts.PostCount)
}

func (p Posts) Count(minWordCount int) int {
	counts := p.Counts(minWordCount)

	return counts.PostCount
}

func (p Posts) Counts(minWordCount int) PostsCounts {
	counts := PostsCounts{}

	for _, post := range p {
		postWordCount := post.WordCount()
		if postWordCount < minWordCount {
			continue
		}

		counts.WordCount += postWordCount
		counts.PostCount++
	}

	return counts
}

func (p Posts) Print() {
	slices.Reverse(p)

	for _, post := range p {
		post.Print()
	}
}
