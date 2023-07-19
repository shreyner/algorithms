package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LessonTime struct {
	Hour   int
	Minute int
}

func NewLessonTimeFromString(str string) *LessonTime {
	values := strings.Split(str, ".")

	lessTime := LessonTime{}

	hour, _ := strconv.Atoi(values[0])
	lessTime.Hour = hour

	if len(values) > 1 {
		minute, _ := strconv.Atoi(values[1])
		lessTime.Minute = minute
	}

	return &lessTime
}

func (l *LessonTime) Less(b *LessonTime) int {
	if l.Hour != b.Hour {
		return l.Hour - b.Hour
	}

	return l.Minute - b.Minute
}

func (l *LessonTime) String() string {
	if l.Minute != 0 {
		return fmt.Sprintf("%v.%v", l.Hour, l.Minute)
	}

	return fmt.Sprintf("%v", l.Hour)
}

type Lesson struct {
	Start  *LessonTime
	Finish *LessonTime
}

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	lessons := make([]*Lesson, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		lesson := Lesson{
			Start:  NewLessonTimeFromString(values[0]),
			Finish: NewLessonTimeFromString(values[1]),
		}

		lessons[i] = &lesson
	}

	sort.Slice(lessons, func(i, j int) bool {
		hLess := lessons[i].Finish.Less(lessons[j].Finish)

		if hLess == 0 {
			return lessons[i].Start.Less(lessons[j].Start) < 0
		}

		return hLess < 0
	})

	results := make([]*Lesson, 0, n)
	results = append(results, lessons[0])

	for i := 1; i < n; i++ {
		if results[len(results)-1].Finish.Less(lessons[i].Start) <= 0 {
			results = append(results, lessons[i])
		}
	}

	fmt.Fprintln(writer, len(results))
	for _, result := range results {
		fmt.Fprintf(writer, "%v %v\n", result.Start.String(), result.Finish.String())
	}
}
