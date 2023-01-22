package neighbors

import "sort"

// TODO: Refactoring
func Neighbors(matrix [][]int, coordinates []int) (result []int) {
	x, y := coordinates[0], coordinates[1]

	if x+1 < len(matrix) {
		result = append(result, matrix[x+1][y])
	}

	if y-1 >= 0 {
		result = append(result, matrix[x][y-1])
	}

	if y+1 < len(matrix[x]) {
		result = append(result, matrix[x][y+1])
	}

	if x-1 >= 0 {
		result = append(result, matrix[x-1][y])
	}

	sort.Ints(result)

	return

	//for i := coordinates[0] - 1; i <= coordinates[0]+1; i++ {
	//	for j := coordinates[1] - 1; j <= coordinates[1]+1; j++ {
	//		if i == coordinates[0] && j == coordinates[1] {
	//			continue
	//		}
	//
	//		if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
	//			continue
	//		}
	//
	//		result = append(result, matrix[i][j])
	//	}
	//}
}
