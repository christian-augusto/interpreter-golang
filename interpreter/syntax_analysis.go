package interpreter

import (
	"container/list"
	"fmt"
)

type syntaxAnalysis struct {
	allCodes        *list.List
	allSentences    *list.List
	currentSentence *sentence
	previousCode    *code
}

// syntaxAnalysis constructor
func newSyntaxAnalysis() *syntaxAnalysis {
	return &syntaxAnalysis{
		allSentences:    list.New(),
		currentSentence: newSentence(),
		previousCode:    nil,
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

		processed, err = sa.processTypeKeyword(code)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = sa.processLiteralValue(code)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = sa.processAttributionSymbol(code)

		if err != nil {
			return err
		}

		if processed {
			continue
		}
	}

	sa.endSentence()

	for e1 := sa.allSentences.Front(); e1 != nil; e1 = e1.Next() {
		sent := e1.Value.(*sentence)

		for e1 := sent.codes.Front(); e1 != nil; e1 = e1.Next() {
			c := e1.Value.(*code)

			fmt.Println(c.toString())
		}

		fmt.Print("\n\n")
	}

	return err
}

func (sa *syntaxAnalysis) endSentence() {
	if !sa.currentSentence.isEmpty() {
		sa.allSentences.PushBack(sa.currentSentence)
		sa.currentSentence = newSentence()
	}
}

func (sa *syntaxAnalysis) pushCodeBack(code *code) {
	sa.currentSentence.codes.PushBack(code)
	sa.previousCode = code
}

func (sa *syntaxAnalysis) processLiteralValue(code *code) (bool, error) {
	if code.isLiteralValue() {
		if sa.previousCode.isMathOperationSymbol() ||
			sa.previousCode.isAttributionSymbol() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			return true, fmt.Errorf("Literal value \"%v\" can't be after %v", code.value, sa.previousCode.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processIdentifier(code *code) (bool, error) {
	if code.isIdentifier() {
		if sa.previousCode.isTypeKeyword() ||
			sa.previousCode.isMathOperationSymbol() ||
			sa.previousCode.isAttributionSymbol() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			return true, fmt.Errorf("Identifier \"%v\" can't be after %v", code.value, sa.previousCode.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processTypeKeyword(code *code) (bool, error) {
	if code.isTypeKeyword() {
		sa.pushCodeBack(code)

		return true, nil
	}

	return false, nil
}

func (sa *syntaxAnalysis) processMathOperationSymbol(code *code) (bool, error) {
	if code.isMathOperationSymbol() {
		if sa.previousCode.isLiteralValue() ||
			sa.previousCode.isIdentifier() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			return true, fmt.Errorf("Math operation symbol \"%v\" can't be after %v", code.value, sa.previousCode.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processAttributionSymbol(code *code) (bool, error) {
	if code.isAttributionSymbol() {
		if sa.previousCode.isIdentifier() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			return true, fmt.Errorf("Math operation symbol \"%v\" can't be after %v", code.value, sa.previousCode.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processLineBreaker(code *code) (bool, error) {
	if code.isLineBreaker() {
		if !sa.previousCode.isMathOperationSymbol() {
			sa.endSentence()
		}
	}

	return false, nil
}
