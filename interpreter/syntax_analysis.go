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

		processed, err = sa.processLineBreaker(code)

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

		processed, err = sa.processIdentifier(code)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = sa.processMathOperationSymbol(code, false)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = sa.processTypeKeyword(code, false)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = sa.processAttributionSymbol(code, false)

		if err != nil {
			return err
		}

		if processed {
			continue
		}
	}

	err = sa.endSentence()

	if err != nil {
		return err
	}

	fmt.Println("Syntax analysis")

	for e1 := sa.allSentences.Front(); e1 != nil; e1 = e1.Next() {
		sent := e1.Value.(*sentence)

		for e1 := sent.codes.Front(); e1 != nil; e1 = e1.Next() {
			c := e1.Value.(*code)

			fmt.Println(c.toString())
		}

		fmt.Print("\n\n")
	}

	fmt.Print("\n\n\n")

	return err
}

func (sa *syntaxAnalysis) endSentence() error {
	if !sa.currentSentence.isEmpty() {
		var err error = nil
		lastCode := sa.currentSentence.codes.Back().Value.(*code)

		_, err = sa.processMathOperationSymbol(lastCode, true)

		if err != nil {
			return err
		}

		_, err = sa.processTypeKeyword(lastCode, true)

		if err != nil {
			return err
		}

		_, err = sa.processAttributionSymbol(lastCode, true)

		if err != nil {
			return err
		}

		sa.allSentences.PushBack(sa.currentSentence)
		sa.currentSentence = newSentence()
	}

	return nil
}

func (sa *syntaxAnalysis) pushCodeBack(code *code) {
	sa.currentSentence.codes.PushBack(code)
	sa.previousCode = code
}

func (sa *syntaxAnalysis) processLineBreaker(code *code) (bool, error) {
	if code.isLineBreaker() {
		if sa.previousCode == nil ||
			!sa.previousCode.isMathOperationSymbol() {
			sa.endSentence()
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processLiteralValue(code *code) (bool, error) {
	if code.isLiteralValue() {
		if sa.previousCode == nil ||
			sa.previousCode.isMathOperationSymbol() ||
			sa.previousCode.isAttributionSymbol() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			return true, syntaxAnalysisError(sa.previousCode.value, code.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processIdentifier(code *code) (bool, error) {
	if code.isIdentifier() {
		if sa.previousCode == nil ||
			sa.previousCode.isTypeKeyword() ||
			sa.previousCode.isMathOperationSymbol() ||
			sa.previousCode.isAttributionSymbol() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			return true, syntaxAnalysisError(sa.previousCode.value, code.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processMathOperationSymbol(code *code, endingSentence bool) (bool, error) {
	if code.isMathOperationSymbol() {
		if endingSentence {
			return true, syntaxAnalysisErrorEndingCode(code.value)
		} else if sa.previousCode != nil &&
			(sa.previousCode.isLiteralValue() ||
				sa.previousCode.isIdentifier()) {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			previousValue := emptyCodeValue

			if sa.previousCode != nil {
				previousValue = sa.previousCode.value
			}

			return true, syntaxAnalysisError(previousValue, code.value)
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processTypeKeyword(code *code, endingSentence bool) (bool, error) {
	if code.isTypeKeyword() {
		if endingSentence {
			return true, syntaxAnalysisErrorEndingCode(code.value)
		} else if sa.previousCode == nil ||
			code.isTypeKeyword() {
			sa.pushCodeBack(code)

			return true, nil
		}
	}

	return false, nil
}

func (sa *syntaxAnalysis) processAttributionSymbol(code *code, endingSentence bool) (bool, error) {
	if code.isAttributionSymbol() {
		if endingSentence {
			return true, syntaxAnalysisErrorEndingCode(code.value)
		} else if sa.previousCode != nil &&
			sa.previousCode.isIdentifier() {

			sa.pushCodeBack(code)

			return true, nil
		} else {
			previousValue := emptyCodeValue

			if sa.previousCode != nil {
				previousValue = sa.previousCode.value
			}

			return true, syntaxAnalysisError(previousValue, code.value)
		}
	}

	return false, nil
}
