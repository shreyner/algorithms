package seletion_sort

import (
	"reflect"
	"testing"
)

type selectionSortTest struct {
	actual   []int
	expected []int
}

var selectionSortTests = []selectionSortTest{
	{actual: []int{4, 9, 5, 3, 8}, expected: []int{3, 4, 5, 8, 9}},
	{actual: []int{5, -1, -3, 15}, expected: []int{-3, -1, 5, 15}},
	{actual: []int{-1}, expected: []int{-1}},
	{actual: []int{}, expected: []int{}},
}

func TestSelectionSort(t *testing.T) {
	for _, test := range selectionSortTests {
		if output := SelectionSort(test.actual); reflect.DeepEqual(output, test.expected) == false {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
