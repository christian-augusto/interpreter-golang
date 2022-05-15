package interpreter

import "fmt"

func charNotValid(char string, currentLine int) error {
	return fmt.Errorf("\"%v\" is not a valid character at line %v", char, currentLine)
}

func floatNumberSeparatorInvalidPosition(char string, currentLine int) error {
	return fmt.Errorf("float number separator \"%v\" invalid position at line %v", char, currentLine)
}

func lineBreakerInsideString(char string, currentLine int) error {
	return fmt.Errorf("\"%v\" is not a valid character inside string at line %v", char, currentLine)
}

func unexpectedToken(str string, currentLine int) error {
	return fmt.Errorf("unexpected token \"%v\" at line %v", str, currentLine)
}

func attributionSymbolInvalidPosition(char string, currentLine int) error {
	return fmt.Errorf("\"%v\" invalid position to attribution symbol at line %v", char, currentLine)
}

func numberInvalidPosition(char string, currentLine int) error {
	return fmt.Errorf("\"%v\" invalid position to number at line %v", char, currentLine)
}

func identifierCharInvalidPosition(char string, currentLine int) error {
	return fmt.Errorf("\"%v\" invalid position to identifier char at line %v", char, currentLine)
}

func prioritySymbolInvalidPosition(char string, currentLine int) error {
	return fmt.Errorf("Invalid position to priority symbol %v at line %v", char, currentLine)
}

func syntaxAnalysisError(previousValue string, currentValue string) error {
	return fmt.Errorf("Code %v has invalid syntax position after %v", currentValue, previousValue)
}

func syntaxAnalysisErrorEndingCode(currentValue string) error {
	return fmt.Errorf("Code %v can't end sentence", currentValue)
}
