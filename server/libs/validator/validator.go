package libs_validator

func checkLength(str string, minLen int, maxLen int) bool {
	l := len(str)
	if l < minLen || l > maxLen {
		return false
	}
	return true
}
