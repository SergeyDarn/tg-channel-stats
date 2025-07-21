package internal

import (
	"fmt"
)

type Reaction struct {
	Count int
	Emoji string
}

func (r Reaction) ToString() string {
	return fmt.Sprintf("{%d %s} ", r.Count, r.Emoji)
}
