package list_form

import (
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func PadLeftInt(str []int, length int) []int {
	if len(str) == length {
		return str
	}

	buf := make([]int, 0, length)

	for i := 0; i < length-len(str); i++ {
		buf = append(buf, 0)
	}

	buf = append(buf, str...)

	return buf
}

func ListForm(first []int, k int) (res []int) {
	second := strings.Split(strconv.Itoa(k), "")

	firstLength := len(first)
	secondLength := len(second)

	max := Max(firstLength, secondLength)

	first = PadLeftInt(first, max)
	second = PadLeftInt(second, max)

}
