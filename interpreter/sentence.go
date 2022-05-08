package interpreter

import (
	"container/list"
	"interpreter-golang/utils"
)

type sentence struct {
	label string
	codes *list.List
}

func newSentence() *sentence {
	return &sentence{}
}

func (s *sentence) isEmpty() bool {
	return utils.StringIsEmpty(s.label)
}

func (s *sentence) setMathOperation() {
	s.label = mathOperationSentenceLabel
}
