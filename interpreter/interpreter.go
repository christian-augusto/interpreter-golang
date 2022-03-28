package interpreter

type interpreter struct {
	allCode         []rune
	lexicalAnalysis *lexicalAnalysis
	syntaxAnalysis  *syntaxAnalysis
}

func NewInterpreter(allCode string) *interpreter {
	return &interpreter{
		allCode:         []rune(allCode),
		lexicalAnalysis: newLexicalAnalysis(),
		syntaxAnalysis:  newSyntaxAnalysis(),
	}
}

func (c *interpreter) Start() error {
	var err error = nil

	err = c.lexicalAnalysis.Start(c.allCode)

	if err != nil {
		return err
	}

	err = c.syntaxAnalysis.Start(c.lexicalAnalysis.allCodes)

	if err != nil {
		return err
	}

	return nil
}
