package quick_sort

import (
	"reflect"
	"testing"
)

type quickSortTest struct {
	actual   []int
	expected []int
}

var quickSortTests = []quickSortTest{
	{actual: []int{}, expected: []int{}},
	{actual: []int{1}, expected: []int{1}},
	{actual: []int{2, 1}, expected: []int{1, 2}},
	{actual: []int{3, 15, -1, 4, 5, 6, 10, 8, 11}, expected: []int{-1, 3, 4, 5, 6, 8, 10, 11, 15}},
}

func TestQuickSort(t *testing.T) {
	for _, test := range quickSortTests {
		if output := QuickSort(test.actual); reflect.DeepEqual(output, test.expected) == false {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
