package convert_to_ox

import (
	"strconv"
	"strings"
)

func ConvertDecToOx(num int) string {
	var sb []int
	lastNum := num

	for lastNum > 0 {
		sb = append(sb, lastNum%2)
		lastNum = lastNum / 2
	}

	var reverseSb strings.Builder
	for i := len(sb) - 1; 0 <= i; i-- {
		reverseSb.WriteString(strconv.Itoa(sb[i]))
	}

	return reverseSb.String()
}
