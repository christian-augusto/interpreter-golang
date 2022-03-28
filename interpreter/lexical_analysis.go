package interpreter

import (
	"container/list"
	"fmt"
	"strings"
)

type lexicalAnalysis struct {
	allCodes    *list.List
	currentCode *code
}

// lexicalAnalysis constructor
func newLexicalAnalysis() *lexicalAnalysis {
	return &lexicalAnalysis{
		allCodes:    list.New(),
		currentCode: newCode(),
	}
}

// Start lexical analysis
func (la *lexicalAnalysis) Start(allCode []rune) error {
	var err error = nil
	currentLine := uint64(1)

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

	for e1 := la.allCodes.Front(); e1 != nil; e1 = e1.Next() {
		code := e1.Value.(*code)

		fmt.Print(code.toString() + " ")
	}

	return err
}

func (la *lexicalAnalysis) charIsALineBreaker(char string) bool {
	return strings.Contains(LINE_BREAKER, char)
}

func (la *lexicalAnalysis) charIsAWhiteSpaceChar(char string) bool {
	return strings.Contains(WHITE_SPACES_CHARS, char)
}

func (la *lexicalAnalysis) charIsAIdentifierChar(char string) bool {
	return strings.Contains(IDENTIFIER_CHARS, char)
}

func (la *lexicalAnalysis) charIsANumber(char string) bool {
	return strings.Contains(NUMBERS_CHARS, char)
}

func (la *lexicalAnalysis) charIsAMathOperationSymbol(char string) bool {
	return strings.Contains(MATH_OPERATIONS_SYMBOLS, char)
}

func (la *lexicalAnalysis) charIsASymbol(char string) bool {
	return strings.Contains(SYMBOLS, char)
}

func (la *lexicalAnalysis) charIsANumberSignal(char string) bool {
	return strings.Contains(NUMBER_SIGNAL_SYMBOLS, char)
}

func (la *lexicalAnalysis) charIsAFloatNumberSeparator(char string) bool {
	return strings.Contains(FLOAT_NUMBER_SEPARATOR, char)
}

func (la *lexicalAnalysis) charIsAStringDelimiter(char string) bool {
	return strings.Contains(STRING_DELIMITERS, char)
}

func (la *lexicalAnalysis) charIsAEcapeChar(char string) bool {
	return strings.Contains(SCAPE_CHARS, char)
}

func (la *lexicalAnalysis) charIsInDictionary(char string) bool {
	return strings.Contains(LANGUAGE_DICTIONARY, char)
}

func (la *lexicalAnalysis) escapedChar(char string) bool {
	valueR := []rune(la.currentCode.value)

	return la.charIsAEcapeChar(string(valueR[len(valueR)-1]))
}

func (la *lexicalAnalysis) endCode() {
	if !la.currentCode.isEmpty() {
		la.allCodes.PushBack(la.currentCode)
		la.currentCode = newCode()
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
		if la.currentCode.isLiteralValue() {
			if la.currentCode.isNumberType() {
				la.endCode()
			}
		} else if la.currentCode.isAMathOperationSymbol() {
			la.endCode()
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processWhiteSpace(char string, line uint64) (bool, error) {
	if la.charIsAWhiteSpaceChar(char) {
		if la.currentCode.isNumberSignalSymbol() {
			la.currentCode.setMathOperationSymbol(la.currentCode.value, line)
		}

		la.endCode()

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processNumber(char string, line uint64) (bool, error) {
	if la.charIsANumber(char) {
		if la.currentCode.isEmpty() || la.currentCode.isNumberSignalSymbol() {
			la.currentCode.setLiteralValue(char, INT_VALUE_TYPE, line)
		} else if la.currentCode.isLiteralValue() {
			valueType := INT_VALUE_TYPE

			if la.currentCode.valueType == FLOAT_VALUE_TYPE {
				valueType = FLOAT_VALUE_TYPE
			} else if la.currentCode.valueType == DOUBLE_VALUE_TYPE {
				valueType = DOUBLE_VALUE_TYPE
			}

			la.currentCode.setLiteralValue(char, valueType, line)
		}

		return true, nil
	}

	return false, nil
}

func (la *lexicalAnalysis) processSymbol(char string, line uint64) (bool, error) {
	if la.charIsASymbol(char) {
		if la.charIsANumberSignal(char) {
			if la.currentCode.isEmpty() {
				la.currentCode.setNumberSignalSymbol(char, line)
			} else if la.currentCode.isAMathOperationSymbol() || la.currentCode.isNumberSignalSymbol() {
				return true, unexpectedToken(la.currentCode.value+char, line)
			}
		} else if la.charIsAMathOperationSymbol(char) {
			if la.currentCode.isAMathOperationSymbol() || la.currentCode.isNumberSignalSymbol() {
				return true, unexpectedToken(la.currentCode.value+char, line)
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
				return true, floatNumberSeparatorInvalidPosition(char, line)
			}

			la.currentCode.setLiteralValue(char, FLOAT_VALUE_TYPE, line)
		}
	}

	return false, nil
}

func (la *lexicalAnalysis) processIdentifierChar(char string, line uint64) (bool, error) {
	if la.charIsAIdentifierChar(char) {
		if la.currentCode.isEmpty() || la.currentCode.isAIdentifier() {
			la.currentCode.setIdentifier(char, line)

			return true, nil
		}
	}

	return false, nil
}
