package leap

func IsLeapYear(year int) bool {
	divisible4 := year%4 == 0
	divisible100 := year%100 == 0
	if divisible4 && !divisible100 {
		return true
	}

	divisible400 := year%400 == 0
	if divisible400 {
		return true
	}

	return false
}
