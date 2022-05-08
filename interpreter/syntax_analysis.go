package interpreter

import (
	"container/list"
)

type syntaxAnalysis struct {
	allCodes        *list.List
	allSentences    *list.List
	currentSentence *sentence
}

// syntaxAnalysis constructor
func newSyntaxAnalysis() *syntaxAnalysis {
	return &syntaxAnalysis{
		allSentences:    list.New(),
		currentSentence: newSentence(),
	}
}

func (sa *syntaxAnalysis) Start(allCodes *list.List) error {
	var err error = nil

	sa.allCodes = allCodes

	for element := sa.allCodes.Front(); element != nil; element = element.Next() {
		code := element.Value.(*code)
		processed := false

		processed, err = sa.processIdentifier(code)

		if err != nil {
			return err
		}

		if processed {
			continue
		}
	}

	return err
}

func (sa *syntaxAnalysis) endSentence() {
	if !sa.currentSentence.isEmpty() {
		sa.allSentences.PushBack(sa.currentSentence)
		sa.currentSentence = newSentence()
	}
}

func (sa *syntaxAnalysis) processIdentifier(code *code) (bool, error) {
	if code.isIdentifier() {
		if sa.currentSentence.isEmpty() {
			sa.currentSentence.setMathOperation()
		}
	}

	return false, nil
}
