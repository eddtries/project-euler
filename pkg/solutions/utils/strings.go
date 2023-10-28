package utils

func ReverseString(s string) (reversedString string) {
	for _, c := range s {
		reversedString = string(c) + reversedString
	}
	return
}
