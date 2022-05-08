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
	return fmt.Errorf("\"%v\" unexpected token at line %v", str, currentLine)
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
