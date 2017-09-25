package common

func MaxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func IsBetweenNumber(x, min, max int) bool {
	if x >= min && x <= max {
		return true
	}
	return false
}
