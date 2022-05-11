package interpreter

import (
	"container/list"
)

type sentence struct {
	attributionSymbolOnSentence bool
	codes                       *list.List
}

func newSentence() *sentence {
	return &sentence{
		attributionSymbolOnSentence: false,
	}
}

func (s *sentence) isEmpty() bool {
	return s.codes.Len() == 0
}
