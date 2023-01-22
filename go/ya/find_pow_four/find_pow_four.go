package find_pow_four

func FindPowFour(num int) bool {
	for num > 1 {
		if num%4 != 0 {
			return false
		}

		num = num / 4
	}

	return true
}
