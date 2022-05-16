package interpreter

import (
	"container/list"
	"fmt"
)

type syntaxAnalysis struct {
	showLogs                           bool
	allSentences                       *list.List
	currentSentence                    *sentence
	previousCode                       *code
	currentSentenceOpenedPriority      int
	currentSentenceAttributionSymbols  int
	currentSentenceHaveTypeKeyword     bool
	currentSentenceAttrubutionFinished bool
}

// syntaxAnalysis constructor
func newSyntaxAnalysis(showLogs bool) *syntaxAnalysis {
	return &syntaxAnalysis{
		showLogs:                           showLogs,
		allSentences:                       list.New(),
		currentSentence:                    newSentence(),
		previousCode:                       nil,
		currentSentenceOpenedPriority:      0,
		currentSentenceAttributionSymbols:  0,
		currentSentenceHaveTypeKeyword:     false,
		currentSentenceAttrubutionFinished: false,
	}
}

// Start syntax analysis
func (sa *syntaxAnalysis) Start(allCodes *list.List) error {
	var err error = nil

	for element := allCodes.Front(); element != nil; element = element.Next() {
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

		processed, err = sa.processOpenPrioritySymbol(code, false)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = sa.processClosePrioritySymbol(code, false)

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

	if sa.showLogs {
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
	}

	return err
}

func (sa *syntaxAnalysis) endSentence() error {
	if !sa.currentSentence.isEmpty() {
		var err error = nil
		firstCode := sa.currentSentence.codes.Front().Value.(*code)
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

		if sa.currentSentenceOpenedPriority > 0 {
			return syntaxAnalysisPriorityNotClosed(firstCode.line)
		}

		sa.allSentences.PushBack(sa.currentSentence)
		sa.currentSentence = newSentence()
		sa.previousCode = nil
		sa.currentSentenceOpenedPriority = 0
		sa.currentSentenceAttributionSymbols = 0
		sa.currentSentenceHaveTypeKeyword = false
		sa.currentSentenceAttrubutionFinished = false
	}

	return nil
}

func (sa *syntaxAnalysis) pushCodeBack(code *code) {
	sa.currentSentence.codes.PushBack(code)
	sa.previousCode = code
}

func (sa *syntaxAnalysis) attributionIsOpened() bool {
	return sa.currentSentenceHaveTypeKeyword &&
		(sa.currentSentenceAttributionSymbols == 0 || !sa.currentSentenceAttrubutionFinished)
}

func (sa *syntaxAnalysis) sentenceFirstCode() *code {
	c := sa.currentSentence.codes.Front().Value.(*code)

	return c
}

func (sa *syntaxAnalysis) processLineBreaker(code *code) (bool, error) {
	if code.isLineBreaker() {
		if sa.previousCode != nil {
			if sa.previousCode.isTypeKeyword() {
				return true, syntaxAnalysisErrorEndingCode(sa.previousCode)
			} else if sa.attributionIsOpened() {
				return true, syntaxAnalysisInvalidAttribution(sa.sentenceFirstCode().line)
			} else if !sa.previousCode.isMathOperationSymbol() {
				sa.endSentence()
			}
		}

		sa.endSentence()
	}

	return false, nil
}

func (sa *syntaxAnalysis) processLiteralValue(code *code) (bool, error) {
	if code.isLiteralValue() {
		if !sa.attributionIsOpened() {
			if sa.previousCode == nil ||
				sa.previousCode.isMathOperationSymbol() ||
				sa.previousCode.isAttributionSymbol() ||
				sa.previousCode.isOpenPrioritySymbol() {

				sa.pushCodeBack(code)

				sa.currentSentenceAttrubutionFinished = true

				return true, nil
			}
		} else {
			return true, syntaxAnalysisInvalidAttribution(sa.sentenceFirstCode().line)
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}

func (sa *syntaxAnalysis) processIdentifier(code *code) (bool, error) {
	if code.isIdentifier() {
		if sa.previousCode == nil ||
			sa.previousCode.isTypeKeyword() ||
			sa.previousCode.isMathOperationSymbol() ||
			sa.previousCode.isAttributionSymbol() ||
			sa.previousCode.isOpenPrioritySymbol() {

			sa.pushCodeBack(code)

			if sa.currentSentenceAttributionSymbols > 0 {
				sa.currentSentenceAttrubutionFinished = true
			}

			return true, nil
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}

func (sa *syntaxAnalysis) processMathOperationSymbol(code *code, endingSentence bool) (bool, error) {
	if code.isMathOperationSymbol() {
		if !sa.attributionIsOpened() {
			if endingSentence {
				return true, syntaxAnalysisErrorEndingCode(code)
			} else if sa.previousCode != nil &&
				(sa.previousCode.isLiteralValue() ||
					sa.previousCode.isIdentifier() ||
					sa.previousCode.isClosePrioritySymbol()) {

				sa.pushCodeBack(code)

				return true, nil
			}
		} else {
			return true, syntaxAnalysisInvalidAttribution(sa.sentenceFirstCode().line)
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}

func (sa *syntaxAnalysis) processTypeKeyword(code *code, endingSentence bool) (bool, error) {
	if code.isTypeKeyword() {
		if endingSentence {
			return true, syntaxAnalysisErrorEndingCode(code)
		} else if sa.previousCode == nil ||
			code.isTypeKeyword() {

			sa.pushCodeBack(code)

			sa.currentSentenceHaveTypeKeyword = true

			return true, nil
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}

func (sa *syntaxAnalysis) processAttributionSymbol(code *code, endingSentence bool) (bool, error) {
	if code.isAttributionSymbol() {
		if sa.currentSentenceAttributionSymbols > 0 {
			return true, syntaxAnalysisManyAttributionSymbolsInASentence(sa.sentenceFirstCode().line)
		} else if endingSentence {
			return true, syntaxAnalysisErrorEndingCode(code)
		} else if sa.previousCode != nil &&
			(sa.previousCode.isIdentifier() ||
				sa.previousCode.isOpenPrioritySymbol()) {

			sa.pushCodeBack(code)

			sa.currentSentenceAttributionSymbols++

			return true, nil
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}

func (sa *syntaxAnalysis) processOpenPrioritySymbol(code *code, endingSentence bool) (bool, error) {
	if code.isOpenPrioritySymbol() {
		if endingSentence {
			return true, syntaxAnalysisErrorEndingCode(code)
		} else if sa.previousCode == nil ||
			sa.previousCode.isIdentifier() ||
			sa.previousCode.isMathOperationSymbol() ||
			sa.previousCode.isAttributionSymbol() ||
			sa.previousCode.isLineBreaker() {

			sa.pushCodeBack(code)

			sa.currentSentenceOpenedPriority++

			sa.currentSentenceAttrubutionFinished = true

			return true, nil
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}

func (sa *syntaxAnalysis) processClosePrioritySymbol(code *code, endingSentence bool) (bool, error) {
	if code.isClosePrioritySymbol() {
		if sa.currentSentenceOpenedPriority == 0 {
			return true, syntaxAnalysisPriorityNotOpened(code)
		} else if sa.previousCode != nil &&
			sa.previousCode.isLiteralValue() ||
			sa.previousCode.isIdentifier() ||
			sa.previousCode.isOpenPrioritySymbol() {

			sa.pushCodeBack(code)

			sa.currentSentenceOpenedPriority--

			return true, nil
		}

		return true, syntaxAnalysisError(sa.previousCode, code)
	}

	return false, nil
}
