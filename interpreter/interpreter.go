package interpreter

type interpreter struct {
	lexicalAnalysis *lexicalAnalysis
	allCode         []rune
}

func NewInterpreter(allCode string) *interpreter {
	return &interpreter{
		allCode:         []rune(allCode),
		lexicalAnalysis: newLexicalAnalysis(),
	}
}

func (c *interpreter) Start() error {
	err := c.lexicalAnalysis.Start(c.allCode)

	return err
}
