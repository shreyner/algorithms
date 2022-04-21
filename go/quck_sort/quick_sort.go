package quick_sort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := []int{}
	more := []int{}

	for _, item := range arr[1:] {
		if item <= pivot {
			less = append(less, item)
		} else {
			more = append(more, item)
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(more)...)
}
