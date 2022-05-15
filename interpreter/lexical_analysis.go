package interpreter

import (
	"container/list"
	"fmt"
	"strings"
)

type lexicalAnalysis struct {
	showLogs    bool
	allCodes    *list.List
	currentCode *code
}

// lexicalAnalysis constructor
func newLexicalAnalysis(showLogs bool) *lexicalAnalysis {
	return &lexicalAnalysis{
		showLogs:    showLogs,
		allCodes:    list.New(),
		currentCode: newCode(),
	}
}

// Start lexical analysis
func (la *lexicalAnalysis) Start(allCode []rune) error {
	var err error = nil
	currentLine := int(1)

	for _, c := range allCode {
		char := string(c)

		if !la.currentCode.isStringType() {
			if !la.charIsInDictionary(char) {
				return charNotValid(char, currentLine)
			}
		}

		if la.charIsLineBreaker(char) {
			currentLine++
		}

		processed := false

		processed, err = la.processCharInsideString(char, currentLine)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = la.processWhiteSpace(char, currentLine)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = la.processLineBreaker(char, currentLine)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = la.processNumber(char, currentLine)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = la.processSymbol(char, currentLine)

		if err != nil {
			return err
		}

		if processed {
			continue
		}

		processed, err = la.processIdentifierChar(char, currentLine)

		if err != nil {
			return err
		}

		if processed {
			continue
		}
	}

	la.endCode()

	if la.showLogs {
		fmt.Println("Lexical analysis")

		for e1 := la.allCodes.Front(); e1 != nil; e1 = e1.Next() {
			c := e1.Value.(*code)

			fmt.Println(c.toString())
		}

		fmt.Print("\n\n\n")
	}

	return err
}

func (la *lexicalAnalysis) charIsLineBreaker(char string) bool {
	return strings.Contains(lineBreakerChars, char)
}

func (la *lexicalAnalysis) charIsWhiteSpace(char string) bool {
	return strings.Contains(whiteSpacesChars, char)
}

func (la *lexicalAnalysis) charIsIdentifier(char string) bool {
	return strings.Contains(identifierChars, char)
}

func (la *lexicalAnalysis) charIsAttributionSymbol(char string) bool {
	return strings.Contains(attributionSymbol, char)
}

func (la *lexicalAnalysis) charIsOpenPrioritySymbol(char string) bool {
	return openPrioritySymbol == char
}

func (la *lexicalAnalysis) charIsClosePrioritySymbol(char string) bool {
	return closePrioritySymbol == char
}

func (la *lexicalAnalysis) charIsNumber(char string) bool {
	return strings.Contains(numbersChars, char)
}

func (la *lexicalAnalysis) charIsMathOperationSymbol(char string) bool {
	return strings.Contains(mathOperationsSymbols, char)
}

func (la *lexicalAnalysis) charIsSymbol(char string) bool {
	return strings.Contains(symbols, char)
}

func (la *lexicalAnalysis) charIsNumberSignal(char string) bool {
	return strings.Contains(numberSignalSymbols, char)
}

func (la *lexicalAnalysis) charIsFloatNumberSeparator(char string) bool {
	return strings.Contains(floatNumberSeparatorSymbol, char)
}

func (la *lexicalAnalysis) charIsStringDelimiter(char string) bool {
	return strings.Contains(stringDelimiterSymbols, char)
}

func (la *lexicalAnalysis) charIsEcapeChar(char string) bool {
	return strings.Contains(scapeChars, char)
}

func (la *lexicalAnalysis) charIsInDictionary(char string) bool {
	return strings.Contains(languageDictionary, char)
}

func (la *lexicalAnalysis) escapedChar(char string) bool {
	valueR := []rune(la.currentCode.value)

	return la.charIsEcapeChar(string(valueR[len(valueR)-1]))
}

func (la *lexicalAnalysis) endCode() {
	if !la.currentCode.isEmpty() {
		if la.currentCode.isNumberSignalSymbol() {
			la.currentCode.changeToMathOperationSymbol()
		}

		la.allCodes.PushBack(la.currentCode)
		la.currentCode = newCode()
	}
}

func (la *lexicalAnalysis) processCharInsideString(char string, line int) (bool, error) {
	if la.currentCode.isLiteralValue() && la.currentCode.isStringType() {
		if la.charIsStringDelimiter(char) && char == la.currentCode.stringDelimiter {
			if la.escapedChar(char) {
				la.currentCode.setLiteralValue(char, la.currentCode.valueType, line)
			} else {
				return false, nil
			}
		} else {
			if la.charIsLineBreaker(char) {
				return true, lineBreakerInsideString(char, line-1)
			}

			la.currentCode.setLiteralValue(char, la.currentCode.valueType, line)
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processLineBreaker(char string, line int) (bool, error) {
	if la.charIsLineBreaker(char) {
		if la.currentCode.isEmpty() ||
			la.currentCode.isLineBreaker() ||
			la.currentCode.isLiteralValue() ||
			la.currentCode.isMathOperationSymbol() ||
			la.currentCode.isIdentifier() ||
			la.currentCode.isKeyword() ||
			la.currentCode.isAttributionSymbol() ||
			la.currentCode.isPrioritySymbol() {
			la.endCode()
			la.currentCode.setLineBreaker(line - 1)
			la.endCode()
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processWhiteSpace(char string, line int) (bool, error) {
	if la.charIsWhiteSpace(char) {
		if la.currentCode.isNumberSignalSymbol() {
			la.currentCode.setMathOperationSymbol(la.currentCode.value, line)
		}

		la.endCode()

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processNumber(char string, line int) (bool, error) {
	if la.charIsNumber(char) {
		if la.currentCode.isEmpty() || la.currentCode.isNumberSignalSymbol() || la.currentCode.isPrioritySymbol() {
			if la.currentCode.isPrioritySymbol() {
				la.endCode()
			}

			la.currentCode.setLiteralValue(char, intValueType, line)
		} else if la.currentCode.isLiteralValue() {
			valueType := intValueType

			if la.currentCode.valueType == floatValueType {
				valueType = floatValueType
			} else if la.currentCode.valueType == doubleValueType {
				valueType = doubleValueType
			}

			la.currentCode.setLiteralValue(char, valueType, line)
		} else if la.currentCode.isIdentifier() {
			la.currentCode.setIdentifier(char, line)
		} else if la.currentCode.isAttributionSymbol() {
			return true, numberInvalidPosition(char, line)
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processSymbol(char string, line int) (bool, error) {
	if la.charIsSymbol(char) {
		if la.charIsNumberSignal(char) {
			if la.currentCode.isEmpty() {
				la.currentCode.setNumberSignalSymbol(char, line)
			} else if la.currentCode.isMathOperationSymbol() || la.currentCode.isNumberSignalSymbol() ||
				la.currentCode.isLiteralValue() {
				return true, unexpectedToken(la.currentCode.value+char, line)
			}
		} else if la.charIsMathOperationSymbol(char) {
			if la.currentCode.isMathOperationSymbol() || la.currentCode.isNumberSignalSymbol() {
				return true, unexpectedToken(la.currentCode.value+char, line)
			}

			la.endCode()
			la.currentCode.setMathOperationSymbol(char, line)
		} else if la.charIsStringDelimiter(char) {
			if la.currentCode.isEmpty() {
				la.currentCode.setStringDelimiterSymbol(char, line)
			} else if char == la.currentCode.stringDelimiter {
				la.endCode()
			} else {
				la.currentCode.setLiteralValue(char, la.currentCode.valueType, line)
			}
		} else if la.charIsFloatNumberSeparator(char) {
			if !(la.currentCode.isLiteralValue() && la.currentCode.isIntNumberType()) {
				return true, floatNumberSeparatorInvalidPosition(char, line)
			}

			la.currentCode.setLiteralValue(char, floatValueType, line)
		} else if la.charIsAttributionSymbol(char) {
			if la.currentCode.isEmpty() {
				la.currentCode.setAttributionSymbol(char, line)
			} else {
				return true, attributionSymbolInvalidPosition(char, line)
			}
		} else if la.charIsOpenPrioritySymbol(char) {
			la.endCode()
			la.currentCode.setOpenPrioritySymbol(char, line)
			la.endCode()
		} else if la.charIsClosePrioritySymbol(char) {
			la.endCode()
			la.currentCode.setClosePrioritySymbol(char, line)
			la.endCode()
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processIdentifierChar(char string, line int) (bool, error) {
	if la.charIsIdentifier(char) {
		if la.currentCode.isEmpty() ||
			la.currentCode.isIdentifier() ||
			la.currentCode.isKeyword() ||
			la.currentCode.isOpenPrioritySymbol() {
			if la.currentCode.isOpenPrioritySymbol() {
				la.endCode()
			}

			la.currentCode.setIdentifier(char, line)
		} else if la.currentCode.isLiteralValueNumberType() {
			return true, identifierCharInvalidPosition(char, line)
		}

		return true, nil
	}

	return false, nil
}
