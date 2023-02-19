package sleight_of_hand

func SleightOfHand(arr []int, k int) int {
	numbers := make([]int, 10)
	fingerCount := k * 2

	for i := 0; i < len(arr); i++ {
		if arr[i] == -1 {
			continue
		}

		numbers[arr[i]] += 1
	}

	points := 0
	for _, numberCount := range numbers {
		if numberCount == 0 || numberCount > fingerCount {
			continue
		}

		points += 1
	}

	return points

}
