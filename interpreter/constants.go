package interpreter

// value types
const INT_VALUE_TYPE = "int"
const FLOAT_VALUE_TYPE = "float"
const DOUBLE_VALUE_TYPE = "double"
const STRING_VALUE_TYPE = "string"
const BOOLEAN_VALUE_TYPE = "boolean"

// dictionary
const LINE_BREAKER = "\n"
const WHITE_SPACES_CHARS = "\t" + " "
const SCAPE_CHARS = "\\"

const ALPHABET_CHARS = "abcdefghijklsmnopqrstuvwxyz"
const NUMBERS_CHARS = "012345679"

const SUM_SYMBOL = "+"
const SUB_SYMBOL = "-"
const MULT_SYMBOL = "*"
const DIV_SYMBOL = "/"
const MOD_SYMBOL = "%"
const EXP_SYMBOL = "^"

const MATH_OPERATIONS_SYMBOLS = SUM_SYMBOL +
	SUB_SYMBOL +
	MULT_SYMBOL +
	DIV_SYMBOL +
	MOD_SYMBOL +
	EXP_SYMBOL

const FLOAT_NUMBER_SEPARATOR = "."
const STRING_DELIMITERS = "'" + "\""

const POSITIVE_NUMBER_SYMBOL = "+"
const NEGATIVE_NUMBER_SYMBOL = "-"
const NUMBER_SIGNAL_SYMBOLS = POSITIVE_NUMBER_SYMBOL + NEGATIVE_NUMBER_SYMBOL

const SYMBOLS = MATH_OPERATIONS_SYMBOLS +
	FLOAT_NUMBER_SEPARATOR +
	STRING_DELIMITERS +
	NUMBER_SIGNAL_SYMBOLS

const LANGUAGE_DICTIONARY = LINE_BREAKER +
	WHITE_SPACES_CHARS +
	SCAPE_CHARS +
	ALPHABET_CHARS +
	NUMBERS_CHARS +
	SYMBOLS

// code labels
const LITERAL_VALUE_LABEL = "literal_value"
const MATH_OPERATION_SYMBOL_LABEL = "math_operation_symbol"
const NUMBER_SIGNAL_SYMBOL_LABEL = "number_signal_symbol"
