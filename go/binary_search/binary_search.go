package binary_search

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		item := arr[mid]

		if item == target {
			return item
		}

		if item < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
