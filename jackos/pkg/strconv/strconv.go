package strconv

// IntToString .
func IntToString(n int) string {

	lastDigit := n % 10
	c := lastDigit + 48
	if n < 10 {
		return string(rune(c))
	} else {
		return IntToString(n/10) + string(rune(c))
	}
}

// StringToInt .
func StringToInt(s string) int {

	v := 0
	for i := 0; i < len(s); i++ {
		d := s[i] - 48
		v = v*10 + int(d)
	}

	return v
}
