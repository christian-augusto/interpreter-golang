package interpreter

// value types
const intValueType = "int"
const floatValueType = "float"
const doubleValueType = "double"
const stringValueType = "string"
const booleanValueType = "boolean"

// dictionary
const lineBreakerChars = "\n" + "\r\n"
const whiteSpacesChars = "\t" + " "
const scapeChars = "\\"

const identifierChars = "abcdefghijklsmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_$"

const keywords = intValueType +
	floatValueType +
	doubleValueType +
	stringValueType

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

const floatNumberSeparatorSymbol = "."
const stringDelimiterSymbols = "'" + "\""

const positiveNumberSymbol = "+"
const negativeNumberSymbol = "-"
const numberSignalSymbols = positiveNumberSymbol + negativeNumberSymbol

const attributionSymbol = "="

const openPrioritySymbol = "("
const closePrioritySymbol = ")"

const prioritySymbols = openPrioritySymbol +
	closePrioritySymbol

const symbols = mathOperationsSymbols +
	floatNumberSeparatorSymbol +
	stringDelimiterSymbols +
	numberSignalSymbols +
	attributionSymbol +
	prioritySymbols

const languageDictionary = lineBreakerChars +
	whiteSpacesChars +
	scapeChars +
	identifierChars +
	numbersChars +
	symbols

// code labels
const lineBreakerLabel = "line_breaker"
const literalValueLabel = "literal_value"
const mathOperationSymbolLabel = "math_operation_symbol"
const numberSignalSymbolLabel = "number_signal_symbol"
const identifierLabel = "identifier"
const attributionSymbolLabel = "attribution_symbol"
const typeKeywordLabel = "type_keyword"
const openPrioritySymbolLabel = "open_priority_symbol"
const closePrioritySymbolLabel = "close_priority_symbol"

// sentence labels
const mathOperationSentenceLabel = "math_operation"

// empty code value
const emptyCodeValue = "empty"
