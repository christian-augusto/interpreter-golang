package utils

func StringLen(str string) uint64 {
	r := []rune(str)

	return uint64(len(r))
}

func StringIsEmpty(str string) bool {
	return len(str) == 0 && str == ""
}
