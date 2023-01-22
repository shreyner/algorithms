package factorization

func Factorization(n int) (res []int) {
	d := 2

	for d*d <= n {
		if n%d == 0 {
			res = append(res, d)
			n = n / d
		} else {
			d += 1
		}
	}

	if n > 1 {
		res = append(res, n)
	}
	return
}
