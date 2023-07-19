package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LessComparator[T any] func(a, b T) int

func partition[T any](arr []T, left, right int, less LessComparator[T]) int {
	pivot := arr[(left+right)/2]
	i := left
	j := right

	for {
		for less(arr[i], pivot) < 0 {
			i += 1
		}

		for less(pivot, arr[j]) < 0 {
			j -= 1
		}

		if i >= j {
			return j
		}

		arr[j], arr[i] = arr[i], arr[j]
		i += 1
		j -= 1
	}
}

func EffectiveQuickSort[T any](arr []T, left, right int, less LessComparator[T]) {
	if right-left < 1 {
		return
	}

	idxPivot := partition(arr, left, right, less)

	EffectiveQuickSort(arr, left, idxPivot, less)
	EffectiveQuickSort(arr, idxPivot+1, right, less)
}

type User struct {
	Login string
	Tasks int
	Fine  int
}

func taskLessCompare(a, b *User) int {
	if a.Tasks < b.Tasks {
		return 1
	}

	if a.Tasks > b.Tasks {
		return -1
	}

	return 0
}

func fineLessCompare(a, b *User) int {
	if a.Fine < b.Fine {
		return -1
	}

	if a.Fine > b.Fine {
		return 1
	}

	return 0
}

func loginLessCompare(a, b *User) int {
	if a.Login < b.Login {
		return -1
	}

	if a.Login > b.Login {
		return 1
	}

	return 0
}

func userLessCompare(a, b *User) int {
	if v := taskLessCompare(a, b); v != 0 {
		return v
	}

	if v := fineLessCompare(a, b); v != 0 {
		return v
	}

	if v := loginLessCompare(a, b); v != 0 {
		return v
	}

	return 0
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	arr := make([]*User, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		name, tasksStr, fineStr := line[0], line[1], line[2]

		tasks, _ := strconv.Atoi(tasksStr)
		fine, _ := strconv.Atoi(fineStr)

		arr[i] = &User{
			Login: name,
			Tasks: tasks,
			Fine:  fine,
		}
	}

	EffectiveQuickSort(arr, 0, len(arr)-1, userLessCompare)

	for _, user := range arr {
		fmt.Println(user.Login)
	}

	writer.Flush()
}
