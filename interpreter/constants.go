package interpreter

// valueTypes
const intValueType = "int"
const floatValueType = "float"
const doubleValueType = "double"
const stringValueType = "string"

// const booleanValueType = "boolean"

// dictionary
const lineBreaker = "\n"
const whiteSpaces = "\t" + " "
const scapeChars = "\\"

// const alphabetChars = "abcdefghijklsmnopqrstuvwxyz"
const numbers = "012345679"

const sumSymbol = "+"
const subSymbol = "-"
const multSymbol = "*"
const divSymbol = "/"
const modSymbol = "%"
const expSymbol = "^"

const mathOperationsSymbols = sumSymbol +
	subSymbol +
	multSymbol +
	divSymbol +
	modSymbol +
	expSymbol

const floatNumberSeparators = "."
const stringDelimiters = "'\""
const numberSignals = "+-"

// const attrSymbol = "="
const symbols = mathOperationsSymbols +
	floatNumberSeparators +
	stringDelimiters +
	numberSignals
	// attrSymbol

const languageDictionary = lineBreaker +
	whiteSpaces +
	// alphabetChars +
	numbers +
	symbols

// code labels
const literalValue = "literal_value"
const mathOperationSymbol = "math_operation_symbol"
const numberSignalSymbol = "number_signal_symbol"

// sentence labels
// const mathOperation = "math_operation"
