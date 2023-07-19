package main

func siftUp(heap []int, idx int) int {
	if idx == 1 {
		return idx
	}

	parentIdx := idx / 2

	if heap[parentIdx] < heap[idx] {
		heap[parentIdx], heap[idx] = heap[idx], heap[parentIdx]

		return siftUp(heap, parentIdx)
	}

	return idx
}

func test() {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}

	if siftUp(sample, 5) != 1 {
		panic("WA")
	}
}

//func main() {
//	test()
//}
