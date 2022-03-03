package interpreter

import "container/list"

type sentence struct {
	// label string
	codes *list.List
}

func newSentence(codes *list.List) *sentence {
	return &sentence{
		codes: codes,
	}
}

// func (s *sentence) setMathOperationSentence() {
// 	s.label = mathOperation
// }
