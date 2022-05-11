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

const floatNumberSeparatorSymbols = "."
const stringDelimiterSymbols = "'" + "\""

const positiveNumberSymbol = "+"
const negativeNumberSymbol = "-"
const numberSignalSymbols = positiveNumberSymbol + negativeNumberSymbol

const attributionSymbols = "="

const symbols = mathOperationsSymbols +
	floatNumberSeparatorSymbols +
	stringDelimiterSymbols +
	numberSignalSymbols +
	attributionSymbols

const languageDictionary = lineBreakerChars +
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
const attributionSymbolLabel = "attribution_symbol"
const typeKeywordLabel = "type_keyword"

// sentence labels
const mathOperationSentenceLabel = "math_operation"
