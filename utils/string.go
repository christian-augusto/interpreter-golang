package utils

func StringLen(str string) uint64 {
	r := []rune(str)

	return uint64(len(r))
}
