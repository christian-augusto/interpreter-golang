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
	line            uint64
}

func newCode() *code {
	return &code{}
}

func (c *code) setLiteralValue(value string, valueType string, line uint64) {
	c.label = LITERAL_VALUE_LABEL
	c.value += value
	c.valueType = valueType
	c.line = line
}

func (c *code) setMathOperationSymbol(value string, line uint64) {
	c.label = MATH_OPERATION_SYMBOL_LABEL
	c.value = value
	c.line = line
}

func (c *code) setNumberSignalSymbol(value string, line uint64) {
	c.label = NUMBER_SIGNAL_SYMBOL_LABEL
	c.value = value
	c.line = line
}

func (c *code) setStringDelimiterSymbol(stringDelimiter string, line uint64) {
	c.label = LITERAL_VALUE_LABEL
	c.valueType = STRING_VALUE_TYPE
	c.stringDelimiter = stringDelimiter
	c.line = line
}

func (c *code) setIdentifier(value string, line uint64) {
	c.label = IDENTIFIER_LABEL
	c.value += value
	c.line = line
}

func (c *code) isEmpty() bool {
	return utils.StringIsEmpty(c.label)
}

func (c *code) isLiteralValue() bool {
	return c.label == LITERAL_VALUE_LABEL
}

func (c *code) isNumberType() bool {
	return c.label == LITERAL_VALUE_LABEL &&
		(c.valueType == INT_VALUE_TYPE ||
			c.valueType == FLOAT_VALUE_TYPE ||
			c.valueType == DOUBLE_VALUE_TYPE)
}

func (c *code) isIntNumberType() bool {
	return c.valueType == INT_VALUE_TYPE
}

// func (c *code) isFloatingPointNumberType() bool {
// 	return c.valueType == FLOAT_VALUE_TYPE ||
// 		c.valueType == DOUBLE_VALUE_TYPE
// }

func (c *code) isStringType() bool {
	return c.valueType == STRING_VALUE_TYPE
}

func (c *code) isAMathOperationSymbol() bool {
	return c.label == MATH_OPERATION_SYMBOL_LABEL
}

func (c *code) isNumberSignalSymbol() bool {
	return c.label == NUMBER_SIGNAL_SYMBOL_LABEL
}

func (c *code) isAIdentifier() bool {
	return c.label == IDENTIFIER_LABEL
}

func (c *code) toString() string {
	mask := `
		{ label: "%v", value: "%v", valueType: "%v", stringDelimiter: "%v", line: %v}
	`

	return fmt.Sprintf(mask, c.label, c.value, c.valueType, c.stringDelimiter, c.line)
}
