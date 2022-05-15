package interpreter

import (
	"fmt"
	"interpreter-golang/utils"
)

type code struct {
	label           string
	value           string
	valueType       string
	stringDelimiter string
	line            int
}

func newCode() *code {
	return &code{}
}

func (c *code) setLineBreaker(line int) {
	c.label = lineBreakerLabel
	c.line = line
}

func (c *code) setLiteralValue(value string, valueType string, line int) {
	c.label = literalValueLabel
	c.value += value
	c.valueType = valueType
	c.line = line
}

func (c *code) setMathOperationSymbol(value string, line int) {
	c.changeToMathOperationSymbol()
	c.value = value
	c.line = line
}

func (c *code) changeToMathOperationSymbol() {
	c.label = mathOperationSymbolLabel
}

func (c *code) setNumberSignalSymbol(value string, line int) {
	c.label = numberSignalSymbolLabel
	c.value = value
	c.line = line
}

func (c *code) setStringDelimiterSymbol(stringDelimiter string, line int) {
	c.label = literalValueLabel
	c.valueType = stringValueType
	c.stringDelimiter = stringDelimiter
	c.line = line
}

func (c *code) setAttributionSymbol(value string, line int) {
	c.label = attributionSymbolLabel
	c.value = value
	c.line = line
}

func (c *code) setIdentifier(value string, line int) {
	c.value += value

	if label := c.getIdentifierKeywordLabel(); label != "" {
		c.label = label
	} else {
		c.label = identifierLabel
	}

	c.line = line
}

func (c *code) setOpenPrioritySymbol(value string, line int) {
	c.label = openPrioritySymbolLabel
	c.value = value
	c.line = line
}

func (c *code) setClosePrioritySymbol(value string, line int) {
	c.label = closePrioritySymbolLabel
	c.value = value
	c.line = line
}

func (c *code) isEmpty() bool {
	return utils.StringIsEmpty(c.label)
}

func (c *code) isLineBreaker() bool {
	return c.label == lineBreakerLabel
}

func (c *code) isLiteralValue() bool {
	return c.label == literalValueLabel
}

func (c *code) isLiteralValueNumberType() bool {
	return c.isLiteralValue() &&
		(c.valueType == intValueType ||
			c.valueType == floatValueType ||
			c.valueType == doubleValueType)
}

func (c *code) isIntNumberType() bool {
	return c.valueType == intValueType
}

func (c *code) isStringType() bool {
	return c.valueType == stringValueType
}

func (c *code) isMathOperationSymbol() bool {
	return c.label == mathOperationSymbolLabel
}

func (c *code) isNumberSignalSymbol() bool {
	return c.label == numberSignalSymbolLabel
}

func (c *code) isAttributionSymbol() bool {
	return c.label == attributionSymbolLabel
}

func (c *code) isPrioritySymbol() bool {
	return c.isOpenPrioritySymbol() || c.isClosePrioritySymbol()
}

func (c *code) isOpenPrioritySymbol() bool {
	return c.label == openPrioritySymbolLabel
}

func (c *code) isClosePrioritySymbol() bool {
	return c.label == closePrioritySymbolLabel
}

func (c *code) isIdentifier() bool {
	return c.label == identifierLabel
}

func (c *code) isKeyword() bool {
	return c.isTypeKeyword()
}

func (c *code) isTypeKeyword() bool {
	return c.label == typeKeywordLabel
}

func (c *code) getIdentifierKeywordLabel() string {
	types := make([]string, 4)

	types[0] = intValueType
	types[1] = floatValueType
	types[2] = doubleValueType
	types[3] = stringValueType

	for _, t := range types {
		if c.value == t {
			return typeKeywordLabel
		}
	}

	return ""
}

func (c *code) toString() string {
	mask := `
		{ label: "%v", value: "%v", valueType: "%v", stringDelimiter: "%v", line: %v}
	`

	return fmt.Sprintf(mask, c.label, c.value, c.valueType, c.stringDelimiter, c.line)
}
