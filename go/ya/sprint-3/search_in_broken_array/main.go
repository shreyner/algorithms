package main

func brokenSearch(arr []int, k int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == k {
			return mid
		}

		if arr[left] <= arr[mid] {
			if arr[left] <= k && k < arr[mid] {
				right = mid
			} else {
				left = mid + 1
			}

			continue
		}

		if arr[mid] < k && k <= arr[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}

func Test() {
	arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}

	if brokenSearch(arr, 5) != 6 {
		panic("WA")
	}
}
