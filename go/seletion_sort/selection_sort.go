package seletion_sort

func findIndexMin(arr []int) int {
	indexMin := 0

	for index, item := range arr {
		if item < arr[indexMin] {
			indexMin = index
		}
	}

	return indexMin
}

func SelectionSort(arr []int) []int {
	length := len(arr)
	sortedArr := make([]int, length)

	for i := 0; i < length; i++ {
		indexMin := findIndexMin(arr)
		sortedArr[i] = arr[indexMin]

		arr = append(arr[:indexMin], arr[indexMin+1:]...)
	}

	return sortedArr
}
