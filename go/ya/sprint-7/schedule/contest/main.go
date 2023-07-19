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

	time := LessonTime{}

	hour, _ := strconv.Atoi(values[0])
	time.Hour = hour

	if len(values) > 1 {
		minute, _ := strconv.Atoi(values[1])
		time.Minute = minute
	}

	return &time
}

func (t *LessonTime) Less(moreTime *LessonTime) int {
	if t.Hour != moreTime.Hour {
		return t.Hour - moreTime.Hour
	}

	return t.Minute - moreTime.Minute
}

func (t LessonTime) String() string {
	if t.Minute > 0 {
		return fmt.Sprintf("%v.%v", t.Hour, t.Minute)
	}

	return fmt.Sprintf("%v", t.Hour)
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

	lessons := make([]*Lesson, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		schedule := strings.Split(scanner.Text(), " ")
		start, finish := schedule[0], schedule[1]

		lesson := Lesson{
			Start:  NewLessonTimeFromString(start),
			Finish: NewLessonTimeFromString(finish),
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

	results := make([]*Lesson, 0, 8)

	results = append(results, lessons[0])

	for i := 1; i < len(lessons); i++ {
		if results[len(results)-1].Finish.Less(lessons[i].Start) <= 0 {
			results = append(results, lessons[i])
		}
	}

	fmt.Fprintln(writer, len(results))

	for _, result := range results {
		fmt.Fprintf(writer, "%v %v\n", result.Start.String(), result.Finish.String())
	}

	writer.Flush()
}
