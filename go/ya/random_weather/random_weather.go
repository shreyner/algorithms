package random_weather

func RandomWeather(arrWeather []int) (count int) {
	if len(arrWeather) < 2 {
		return len(arrWeather)
	}

	for i, v := range arrWeather {
		if i == 0 {
			if v > arrWeather[i+1] {
				count++
			}

			continue
		}

		if i == len(arrWeather)-1 {
			if v > arrWeather[i-1] {
				count++
			}
			continue
		}

		if arrWeather[i-1] < v && v > arrWeather[i+1] {
			count++
			continue
		}
	}

	return
}

//func main() {
//	var numRows int
//	fmt.Scan(&numRows)
//
//	reader := bufio.NewReader(os.Stdin)
//	scanner := bufio.NewScanner(reader)
//	scanner.Split(bufio.ScanLines)
//	buf := make([]byte, 0, 64*1024)
//	scanner.Buffer(buf, 1024*1024)
//
//	writer := bufio.NewWriter(os.Stdout)
//
//	nums := make([]int, numRows)
//
//	b := scanner.Scan()
//	if !b {
//		fmt.Println("Error", scanner.Err())
//		return
//	}
//
//	line := scanner.Text()
//	values := strings.Split(line, " ")
//
//	for i := 0; i < numRows; i++ {
//		num, _ := strconv.Atoi(values[i])
//
//		nums[i] = num
//	}
//
//	result := RandomWeather(nums)
//
//	output_string := strconv.Itoa(result)
//
//	writer.WriteString(output_string)
//	writer.WriteString("\n")
//	writer.Flush()
//}
