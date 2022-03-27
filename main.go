package main

import (
	"fmt"
	"interpreter-golang/dependencies"
	"interpreter-golang/interpreter"
	"interpreter-golang/logger"
)

func main() {
	var logger dependencies.Logger = logger.NewLogger()

	test1(logger)
	// test2(logger)
	// test3(logger)
	// test4(logger)
	// test5(logger)
	// test6(logger)
	// test7(logger)
	// test8(logger)
	// test9(logger)
	// test10(logger)
	// test11(logger)
	// test12(logger)
	// test13(logger)
	// test14(logger)
}

func test1(logger dependencies.Logger) {
	code := `
		1 + 2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test2(logger dependencies.Logger) {
	code := `
		"a+" + "a"
		2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test3(logger dependencies.Logger) {
	code := `
		"'"
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test4(logger dependencies.Logger) {
	code := `
		"\""
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test5(logger dependencies.Logger) {
	code := `
		2
		"
		"
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test6(logger dependencies.Logger) {
	code := `
		.
		2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test7(logger dependencies.Logger) {
	code := `
		1 + 1
		2.222 - 2
		3 * 3
		4 / 4
		5 % 5
		6 ^ 6
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test8(logger dependencies.Logger) {
	code := `
		5.555.
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test9(logger dependencies.Logger) {
	code := `
		1 - -2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test10(logger dependencies.Logger) {
	code := `
		1 - -2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test11(logger dependencies.Logger) {
	code := `
		1 *-2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test12(logger dependencies.Logger) {
	code := `
		1 * -2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test13(logger dependencies.Logger) {
	code := `
		1 + "2"
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}

func test14(logger dependencies.Logger) {
	code := `
		a 2
	`

	comp := interpreter.NewInterpreter(code)

	err := comp.Start()

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println("------------")
}
