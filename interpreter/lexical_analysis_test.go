package interpreter

import (
	"testing"
)

func TestStart(t *testing.T) {
	case1(t)
	case2(t)
	case3(t)
	case4(t)
	case5(t)
	case6(t)
	case7(t)
	case8(t)
	case9(t)
	case10(t)
	case11(t)
	case12(t)
	case13(t)
	case14(t)
	case15(t)
	case16(t)
	case17(t)
	case18(t)
	case19(t)
}

func case1(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		1 + 2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case2(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		"a+" + "a"
		2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "a+" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "a" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case3(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		"'"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "'" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}
	}
}

func case4(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		"\""
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "\\\"" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}
	}
}

func case5(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

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

func case6(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		.
		2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Float delimiter in invalid position")
	}
}

func case7(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

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
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2.222" || c.valueType != "float" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[3] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "-" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[4] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[5] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "3" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[6] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "*" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[7] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "3" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 4 {
			t.Errorf("Code[8] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "4" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[9] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "/" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[10] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "4" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 5 {
			t.Errorf("Code[11] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "5" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[12] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "%" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[11] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "5" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 6 {
			t.Errorf("Code[12] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "6" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[13] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "^" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[14] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "6" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 7 {
			t.Errorf("Code[15] isn't in the current value %v", c)
		}
	}
}

func case8(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		5.555.
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Float delimiter in invalid position")
	}
}

func case9(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		1 - -2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "-" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "-2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case10(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		1 *-2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token *-")
	}
}

func case11(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		1 * -2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "*" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "-2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case12(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		1 + "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "literal_value" || c.value != "1" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "math_operation_symbol" || c.value != "+" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case13(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		a2 = 2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "identifier" || c.value != "a2" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "int" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case14(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		a2= 2
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token a2=")
	}
}

func case15(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		a = "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "identifier" || c.value != "a" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 2 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case16(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		a
		= "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "identifier" || c.value != "a" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 3 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case17(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		a2
		= "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "identifier" || c.value != "a2" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 3 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}

func case18(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		2a
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err == nil {
		t.Error("Invalid token 2a")
	}
}

func case19(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis()

	codeStr = `
		int
		= "2"
	`

	err = lexicalAnalysis.Start([]rune(codeStr))

	if err != nil {
		t.Error(err)
	} else {
		elem := lexicalAnalysis.allCodes.Front()
		c := elem.Value.(*code)

		if c.label != "keyword" || c.value != "int" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 2 {
			t.Errorf("Code[0] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "attribution_symbol" || c.value != "=" || c.valueType != "" || c.stringDelimiter != "" ||
			c.line != 3 {
			t.Errorf("Code[1] isn't in the current value %v", c)
		}

		elem = elem.Next()
		c = elem.Value.(*code)

		if c.label != "literal_value" || c.value != "2" || c.valueType != "string" || c.stringDelimiter != "\"" ||
			c.line != 3 {
			t.Errorf("Code[2] isn't in the current value %v", c)
		}
	}
}
