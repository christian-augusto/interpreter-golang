package interpreter

import "container/list"

type sentence struct {
	codes *list.List
}

func newSentence(codes *list.List) *sentence {
	return &sentence{
		codes: codes,
	}
}
