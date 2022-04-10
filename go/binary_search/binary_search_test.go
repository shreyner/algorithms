package binary_search

import "testing"

type binarySearchTest struct {
	arr      []int
	target   int
	expected int
}

var binarySearchTests = []binarySearchTest{
	{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 7, expected: 7},
	{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 2, expected: 2},
	{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 9, expected: 9},
	{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 1, expected: 1},
	{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 15, expected: -1},
	{arr: []int{2, 3, 4, 5, 6, 7, 8, 9}, target: 1, expected: -1},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range binarySearchTests {
		if output := BinarySearch(test.arr, test.target); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
