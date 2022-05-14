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

	showLogs := false

	if len(os.Args) > 2 && os.Args[2] == "showLogs" {
		showLogs = true
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

	inter := interpreter.NewInterpreter(code, showLogs)

	err = inter.Start()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
