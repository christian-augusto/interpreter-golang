package utils

func StringLen(str string) int {
	r := []rune(str)

	return int(len(r))
}

func StringIsEmpty(str string) bool {
	return len(str) == 0 && str == ""
}
