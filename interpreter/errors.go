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
