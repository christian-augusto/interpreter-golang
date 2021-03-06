package interpreter

import (
	"testing"
)

func TestSyntaxAnalysis1(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `1 + 2`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 3 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis2(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `1 + 2 +`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Code + can't end sentence")
	}
}

func TestSyntaxAnalysis3(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `1 + 2 int`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Code int (type keyword) can't end sentence")
	}
}

func TestSyntaxAnalysis4(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `1 + 2 =`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Code = (attribution symbol) can't end sentence")
	}
}

func TestSyntaxAnalysis5(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int 1`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Code 1 (literal value) has invalid syntax position after int (type keyword)")
	}
}

func TestSyntaxAnalysis6(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int i`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 2 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis7(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int i = a + 2 - 3 * 4`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 10 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis8(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Code int (type keyword) can't end sentence")
	}
}

func TestSyntaxAnalysis9(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int a = (3 + 3)`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 8 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis10(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int a = (3 + )`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Code ) (close_priority_symbol) has invalid syntax position after + (math_operation_symbol)")
	}
}

func TestSyntaxAnalysis11(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `a(3 + 3)`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 6 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis12(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `
		int
		a
	`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Invalid attribution syntax sentence")
	}
}

func TestSyntaxAnalysis13(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int a + a = 2`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Invalid attribution syntax sentence")
	}
}

func TestSyntaxAnalysis14(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `
		int a
		= 2
	`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Invalid attribution syntax sentence")
	}
}

func TestSyntaxAnalysis15(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `
		int a =
		2
	`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Invalid attribution syntax sentence")
	}
}

func TestSyntaxAnalysis16(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int a = b`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 4 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis17(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `int a = a = b`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err == nil {
		t.Errorf("Invalid attribution syntax sentence count in a sentece")
	}
}

func TestSyntaxAnalysis18(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `string a`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 2 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}

func TestSyntaxAnalysis19(t *testing.T) {
	var err error
	var codeStr string
	lexicalAnalysis := newLexicalAnalysis(false)
	syntaxAnalysis := newSyntaxAnalysis(false)

	codeStr = `string b = "2"`

	lexicalAnalysis.Start([]rune(codeStr))
	err = syntaxAnalysis.Start(lexicalAnalysis.allCodes)

	if err != nil {
		t.Error(err)
	} else {
		if syntaxAnalysis.allSentences.Len() != 1 {
			t.Errorf("syntaxAnalysis.allSentences.Len() invalid value %v", syntaxAnalysis.allSentences.Len())
			return
		}

		sentence := syntaxAnalysis.allSentences.Front().Value.(*sentence)

		if sentence.codes.Len() != 4 {
			t.Errorf("sentence.codes.Len() invalid value %v", sentence.codes.Len())
		}
	}
}
