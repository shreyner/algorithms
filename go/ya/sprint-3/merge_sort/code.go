package merge_sort

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf < 2 {
		return
	}

	mid := (lf + rg) / 2

	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)

	res := merge(arr, lf, mid, rg)

	for i := lf; i < rg; i++ {
		arr[i] = res[i-lf]
	}
}

func merge(arr []int, left int, mid int, right int) (result []int) {
	//result := make([]int, right)
	l, r := left, mid

	for l < mid && r < right {
		if arr[l] <= arr[r] {
			result = append(result, arr[l])
			l += 1
		} else {
			result = append(result, arr[r])
			r += 1
		}
	}

	for l < mid {
		result = append(result, arr[l])
		l += 1
	}

	for r < right {
		result = append(result, arr[r])
		r += 1
	}

	return result
}

//func main() {
//	a := []int{1, 4, 9, 2, 10, 11}
//	b := merge(a, 0, 3, 6)
//	expected := []int{1, 2, 4, 9, 10, 11}
//	fmt.Println(b, expected)
//	if !reflect.DeepEqual(b, expected) {
//		panic("WA. Merge")
//	}
//
//	c := []int{1, 4, 2, 10, 1, 2}
//	merge_sort(c, 0, 6)
//	expected = []int{1, 1, 2, 2, 4, 10}
//	fmt.Println(c, expected)
//	if !reflect.DeepEqual(c, expected) {
//		panic("WA. MergeSort")
//	}
//}
