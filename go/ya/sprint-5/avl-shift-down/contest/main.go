package main

func siftDown(heap []int, idx int) int {
	leftIdx := 2 * idx
	rightIds := 2*idx + 1

	if len(heap)-1 < leftIdx {
		return idx
	}

	var indexLarge int
	if (rightIds <= len(heap)-1) && (heap[leftIdx] < heap[rightIds]) {
		indexLarge = rightIds
	} else {
		indexLarge = leftIdx
	}

	if heap[idx] < heap[indexLarge] {
		heap[indexLarge], heap[idx] = heap[idx], heap[indexLarge]

		return siftDown(heap, indexLarge)
	}

	return idx
}

func test() {
	//sample := []int{-1, 12, 1, 8, 3, 4, 7}
	//
	//if siftDown(sample, 2) != 5 {
	//	panic("WA")
	//}

	//sample := []int{3, 3, 2, 1}
	//
	//fmt.Println(siftDown(sample, 2))
}

//func main() {
//	test()
//}
