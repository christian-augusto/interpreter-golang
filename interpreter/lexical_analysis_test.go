package interpreter

import (
	"testing"
)

func TestLexicalAnalysis1(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `1 + 2`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis2(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		"a+" + "a"
		2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 7 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "a+" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "a" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[3] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[4] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[5] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[6] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis3(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `"'"`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 1 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "'" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis4(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `"\""`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 1 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "\\\"" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis5(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		2
		"
		"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Line breaker isn't allowed inside \"\" delimiters")
	}
}

func TestLexicalAnalysis6(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		.
		2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Float delimiter in invalid position")
	}
}

func TestLexicalAnalysis7(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		1 + 1
		2.222 - 2
		3 * 3
		4 / 4
		5 % 5
		6 ^ 6
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error("Float delimiter in invalid position")
	} else {
		if lexicalAnalysis.allCodes.Len() != 25 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[3] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[4] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2.222" || c.valueType != "float" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[5] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "-" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[6] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[7] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[8] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "3" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[9] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "*" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[10] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "3" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[11] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[12] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "4" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[13] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "/" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[14] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "4" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[15] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[16] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "5" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[17] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "%" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[18] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "5" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[19] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[20] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "6" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[21] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "^" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[22] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "6" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[23] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[24] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis8(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `5.555.`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Float delimiter in invalid position")
	}
}

func TestLexicalAnalysis9(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `1 - -2`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "-" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "-2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis10(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `1 *-2`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token *-")
	}
}

func TestLexicalAnalysis11(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `1 * -2`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "*" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "-2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis12(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `1 + "2"`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis13(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `a2 = 2`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "identifier" || c.value != "a2" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis14(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `a2= 2`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token a2=")
	}
}

func TestLexicalAnalysis15(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `a = "2"`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "identifier" || c.value != "a" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis16(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		a
		= "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 6 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "identifier" || c.value != "a" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[3] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 3 {
			t.Errorf("Code[4] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[5] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis17(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		a2
		= "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 6 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "identifier" || c.value != "a2" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[3] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 3 {
			t.Errorf("Code[4] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[5] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis18(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `2a`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token 2a")
	}
}

func TestLexicalAnalysis19(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `
		int
		= "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 6 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "type_keyword" || c.value != "int" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[3] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 3 {
			t.Errorf("Code[4] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[5] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis20(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `+1+1`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token +1+")
	}
}

func TestLexicalAnalysis21(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `


	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis22(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `1 +`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis23(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `()`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis24(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `(`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 1 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis25(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 1 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis26(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `(a)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "identifier" || c.value != "a" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis27(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `(1)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis28(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `(
		1
	)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 5 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "line_breaker" || c.value != "" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[3] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[4] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis29(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `(+)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 3 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis30(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `+(`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis31(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `2(`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis32(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `"2"(`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis33(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `2)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis34(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `"2")`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis35(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `=(`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "open_priority_symbol" || c.value != "(" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}

func TestLexicalAnalysis36(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)

	codeStr = `=)`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		if lexicalAnalysis.allCodes.Len() != 2 {
			t.Errorf("lexicalAnalysis.allCodes.Len() invalid value %v", lexicalAnalysis.allCodes.Len())
			return
		}

		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[0] isn't in the current value %v", c)
			return
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "close_priority_symbol" || c.value != ")" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 1 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}
	}
}
