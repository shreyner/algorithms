package nearest_zero

func NearestZero(arr []int) []int {
	resultArr := make([]int, len(arr))

	count := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count = 0
		}

		if count == -1 {
			continue
		}

		resultArr[i] = count

		count += 1
	}

	count = -1

	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] == 0 {
			count = 0
		}

		if count == -1 {
			continue
		}

		if resultArr[j] == 0 || resultArr[j] > count {
			resultArr[j] = count
		}

		count += 1
	}

	return resultArr
}
