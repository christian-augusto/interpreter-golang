package interpreter

import "fmt"

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
	c.label = literalValue
	c.value += value
	c.valueType = valueType
	c.line = line
}

func (c *code) setMathOperationSymbol(value string, line uint64) {
	c.label = mathOperationSymbol
	c.value = value
	c.line = line
}

func (c *code) setNumberSignalSymbol(value string, line uint64) {
	c.label = numberSignalSymbol
	c.value = value
	c.line = line
}

func (c *code) setStringDelimiterSymbol(stringDelimiter string, line uint64) {
	c.label = literalValue
	c.valueType = stringValueType
	c.stringDelimiter = stringDelimiter
	c.line = line
}

func (c *code) isEmpty() bool {
	return c.label == ""
}

func (c *code) isLiteralValue() bool {
	return c.label == literalValue
}

func (c *code) isNumberType() bool {
	return c.label == literalValue &&
		(c.valueType == intValueType ||
			c.valueType == floatValueType ||
			c.valueType == doubleValueType)
}

func (c *code) isIntNumberType() bool {
	return c.valueType == intValueType
}

func (c *code) isNumberSignalSymbol() bool {
	return c.label == numberSignalSymbol
}

// func (c *code) isFloatingPointNumberType() bool {
// 	return c.valueType == floatValueType ||
// 		c.valueType == doubleValueType
// }

func (c *code) isStringType() bool {
	return c.valueType == stringValueType
}

func (c *code) isAMathOperationSymbol() bool {
	return c.label == mathOperationSymbol
}

func (c *code) toString() string {
	mask := `
		{ label: "%v", value: "%v", valueType: "%v", stringDelimiter: "%v", line: %v}
	`

	return fmt.Sprintf(mask, c.label, c.value, c.valueType, c.stringDelimiter, c.line)
}
