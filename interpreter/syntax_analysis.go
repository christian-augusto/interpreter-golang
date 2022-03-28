package interpreter

import "container/list"

type syntaxAnalysis struct {
	allCodes *list.List
}

// syntaxAnalysis constructor
func newSyntaxAnalysis() *syntaxAnalysis {
	return &syntaxAnalysis{}
}

func (sa *syntaxAnalysis) Start(allCodes *list.List) error {
	var err error = nil

	sa.allCodes = allCodes

	return err
}
