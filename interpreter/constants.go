package interpreter

// value types
const intValueType = "int"
const floatValueType = "float"
const doubleValueType = "double"
const stringValueType = "string"
const booleanValueType = "boolean"

// dictionary
const lineBreaker = "\n"
const whiteSpacesChars = "\t" + " "
const scapeChars = "\\"

const identifierChars = "abcdefghijklsmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_$"
const numbersChars = "012345679"

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

const floatNumberSeparator = "."
const stringDelimiters = "'" + "\""

const positiveNumberSymbol = "+"
const negativeNumberSymbol = "-"
const numberSignalSymbols = positiveNumberSymbol + negativeNumberSymbol

const symbols = mathOperationsSymbols +
	floatNumberSeparator +
	stringDelimiters +
	numberSignalSymbols

const languageDictionary = lineBreaker +
	whiteSpacesChars +
	scapeChars +
	identifierChars +
	numbersChars +
	symbols

// code labels
const literalValueLabel = "literal_value"
const mathOperationSymbolLabel = "math_operation_symbol"
const numberSignalSymbolLabel = "number_signal_symbol"
const identifierLabel = "identifier"
