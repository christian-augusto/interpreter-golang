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
		codes:                       list.New(),
	}
}

func (s *sentence) isEmpty() bool {
	return s.codes.Len() == 0
}
