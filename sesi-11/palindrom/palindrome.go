package palindrom

func IsPalindrom(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
