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

func syntaxAnalysisError(previousCode *code, currentCode *code) error {
	previousCodeValue := "empty"
	previousCodeLabel := "empty"

	if previousCode != nil {
		previousCodeValue = previousCode.value
		previousCodeLabel = previousCode.label
	}

	return fmt.Errorf(
		"Code %v (%v) has invalid syntax position after %v (%v)",
		currentCode.value, currentCode.label, previousCodeValue, previousCodeLabel,
	)
}

func syntaxAnalysisErrorEndingCode(currentCode *code) error {
	return fmt.Errorf(
		"Code %v (%v) can't end sentence at line %v", currentCode.value, currentCode.label, currentCode.line,
	)
}

func syntaxAnalysisInvalidAttribution(line int) error {
	return fmt.Errorf("Invalid attribution syntax sentence started at line %v", line)
}

func syntaxAnalysisManyAttributionSymbolsInASentence(line int) error {
	return fmt.Errorf("Invalid attribution syntax count in a sentece started at line %v", line)
}

func syntaxAnalysisPriorityNotClosed(line int) error {
	return fmt.Errorf("Sentence started at line %v has unclosed priority", line)
}

func syntaxAnalysisPriorityNotOpened(currentCode *code) error {
	return fmt.Errorf(
		"Code %v (%v) is closing a unopened priority at line %v",
		currentCode.value, currentCode.label, currentCode.line,
	)
}
