package math

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func IntMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}