package interpreter

import (
	"container/list"
	"fmt"
	"strings"
)

type lexicalAnalysis struct {
	sentences    *list.List
	currentCodes *list.List
	currentCode  *code
}

// LexicalAnalysis constructor
func newLexicalAnalysis() *lexicalAnalysis {
	return &lexicalAnalysis{
		sentences:    list.New(),
		currentCodes: list.New(),
		currentCode:  newCode(),
	}
}

// Start lexical analysis
func (la *lexicalAnalysis) Start(allCode []rune) error {
	var currentLine uint64 = 1

	for _, c := range allCode {
		char := string(c)

		if !la.currentCode.isStringType() {
			if !la.charIsInDictionary(char) {
				return charNotValid(char, currentLine)
			}
		}

		if la.charIsALineBreaker(char) {
			currentLine++
		}

		var err error = nil

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

		err = la.processSymbol(char, currentLine)

		if err != nil {
			return err
		}
	}

	la.endCode()
	la.endSentence()

	fmt.Println(la.sentences.Len())
	for e1 := la.sentences.Front(); e1 != nil; e1 = e1.Next() {
		fmt.Print("[\n")

		sentence := e1.Value.(*sentence)

		for e2 := sentence.codes.Front(); e2 != nil; e2 = e2.Next() {
			code := e2.Value.(*code)

			fmt.Print(code.toString() + " ")
		}

		fmt.Print("\n]\n")
	}

	return nil
}

func (la *lexicalAnalysis) charIsALineBreaker(char string) bool {
	return strings.Contains(lineBreaker, char)
}

func (la *lexicalAnalysis) charIsAWhiteSpace(char string) bool {
	return strings.Contains(whiteSpaces, char)
}

// func (la *lexicalAnalysis) charIsAAlphabetChar(char string) bool {
// 	return strings.Contains(alphabetChars, char)
// }

func (la *lexicalAnalysis) charIsANumber(char string) bool {
	return strings.Contains(numbers, char)
}

func (la *lexicalAnalysis) charIsAMathOperationSymbol(char string) bool {
	return strings.Contains(mathOperationsSymbols, char)
}

func (la *lexicalAnalysis) charIsASymbol(char string) bool {
	return strings.Contains(symbols, char)
}

func (la *lexicalAnalysis) charIsANumberSignal(char string) bool {
	return strings.Contains(numberSignals, char)
}

func (la *lexicalAnalysis) charIsAFloatNumberSeparator(char string) bool {
	return strings.Contains(floatNumberSeparators, char)
}

func (la *lexicalAnalysis) charIsAStringDelimiter(char string) bool {
	return strings.Contains(stringDelimiters, char)
}

func (la *lexicalAnalysis) charIsAEcapeChar(char string) bool {
	return strings.Contains(scapeChars, char)
}

func (la *lexicalAnalysis) charIsInDictionary(char string) bool {
	return strings.Contains(languageDictionary, char)
}

func (la *lexicalAnalysis) escapedChar(char string) bool {
	valueR := []rune(la.currentCode.value)

	return la.charIsAEcapeChar(string(valueR[len(valueR)-1]))
}

func (la *lexicalAnalysis) endCode() {
	if !la.currentCode.isEmpty() {
		la.currentCodes.PushBack(la.currentCode)
		la.currentCode = newCode()
	}
}

func (la *lexicalAnalysis) endSentence() {
	if la.currentCodes.Len() > 0 {
		la.sentences.PushBack(newSentence(la.currentCodes))
		la.currentCodes = list.New()
	}
}

func (la *lexicalAnalysis) processCharInsideString(char string, line uint64) (bool, error) {
	if la.currentCode.isLiteralValue() && la.currentCode.isStringType() {
		if la.charIsAStringDelimiter(char) && char == la.currentCode.stringDelimiter {
			if la.escapedChar(char) {
				la.currentCode.setLiteralValue(char, la.currentCode.valueType, line)

				return true, nil
			} else {
				return false, nil
			}

		} else {
			if la.charIsALineBreaker(char) {
				return true, lineBreakerInsideString(char, line-1)
			}

			la.currentCode.setLiteralValue(char, la.currentCode.valueType, line)

			return true, nil
		}
	}

	return false, nil
}

func (la *lexicalAnalysis) processLineBreaker(char string, line uint64) (bool, error) {
	if la.charIsALineBreaker(char) {
		if la.currentCode.isEmpty() {
			la.endSentence()
		} else if la.currentCode.isLiteralValue() {
			if la.currentCode.isNumberType() {
				la.endCode()
				la.endSentence()
			}
		} else if la.currentCode.isAMathOperationSymbol() {
			la.endCode()
			la.endSentence()
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processWhiteSpace(char string, line uint64) (bool, error) {
	if la.charIsAWhiteSpace(char) {
		if la.currentCode.isLiteralValue() {
			if la.currentCode.isNumberType() {
				la.endCode()
			}
		} else if la.currentCode.isNumberSignalSymbol() {
			la.currentCode.setMathOperationSymbol(la.currentCode.value, line)
			la.endCode()
		} else if la.currentCode.isAMathOperationSymbol() {
			la.endCode()
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processNumber(char string, line uint64) (bool, error) {
	if la.charIsANumber(char) {
		if la.currentCode.isEmpty() || la.currentCode.isNumberSignalSymbol() {
			la.currentCode.setLiteralValue(char, intValueType, line)
		} else if la.currentCode.label == literalValue {
			valueType := intValueType

			if la.currentCode.valueType == floatValueType {
				valueType = floatValueType
			} else if la.currentCode.valueType == doubleValueType {
				valueType = doubleValueType
			}

			la.currentCode.setLiteralValue(char, valueType, line)
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processSymbol(char string, line uint64) error {
	if la.charIsASymbol(char) {
		if la.charIsANumberSignal(char) {
			if la.currentCode.isEmpty() {
				la.currentCode.setNumberSignalSymbol(char, line)
			} else if la.currentCode.isAMathOperationSymbol() || la.currentCode.isNumberSignalSymbol() {
				return unexpectedToken(la.currentCode.value+char, line)
			}
		} else if la.charIsAMathOperationSymbol(char) {
			if la.currentCode.isAMathOperationSymbol() || la.currentCode.isNumberSignalSymbol() {
				return unexpectedToken(la.currentCode.value+char, line)
			}

			la.endCode()
			la.currentCode.setMathOperationSymbol(char, line)
		} else if la.charIsAStringDelimiter(char) {
			if la.currentCode.isEmpty() {
				la.currentCode.setStringDelimiterSymbol(char, line)
			} else if char == la.currentCode.stringDelimiter {
				la.endCode()
			} else {
				la.currentCode.setLiteralValue(char, la.currentCode.valueType, line)
			}
		} else if la.charIsAFloatNumberSeparator(char) {
			if !(la.currentCode.isLiteralValue() && la.currentCode.isIntNumberType()) {
				return floatNumberSeparatorInvalidPosition(char, line)
			}

			valueType := floatValueType

			if la.currentCode.isStringType() {
				valueType = stringValueType
			}

			la.currentCode.setLiteralValue(char, valueType, line)
		}
	}

	return nil
}
