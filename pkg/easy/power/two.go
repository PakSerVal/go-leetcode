package power

func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}

		n /= 2
	}

	return true
}
