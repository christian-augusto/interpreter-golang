package main

import (
	"fmt"
	"interpreter-golang/dependencies"
	"interpreter-golang/interpreter"
	"interpreter-golang/logger"
	"os"
)

func main() {
	var logger dependencies.Logger = logger.NewLogger()
	var err error

	if len(os.Args) < 2 {
		err = fmt.Errorf("Missing file path to read the code")
		logger.Error(err)
		os.Exit(1)
	}

	filePath := os.Args[1]
	var codeBytes []byte
	var code string

	codeBytes, err = os.ReadFile(filePath)

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	code = string(codeBytes)

	fmt.Println(code)

	inter := interpreter.NewInterpreter(code)

	err = inter.Start()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
