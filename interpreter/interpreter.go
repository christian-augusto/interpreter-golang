package interpreter

type interpreter struct {
	allCode         []rune
	lexicalAnalysis *lexicalAnalysis
	syntaxAnalysis  *syntaxAnalysis
}

func NewInterpreter(allCode string, showLogs bool) *interpreter {
	return &interpreter{
		allCode:         []rune(allCode),
		lexicalAnalysis: newLexicalAnalysis(showLogs),
		syntaxAnalysis:  newSyntaxAnalysis(showLogs),
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

	return err
}
