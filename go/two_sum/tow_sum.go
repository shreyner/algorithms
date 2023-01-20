package two_sum

func TwoSum(nums []int, target int) []int {
	numsLen := len(nums)
	numMap := make(map[int]int, numsLen)

	for i := 0; i < numsLen; i++ {
		value, ok := numMap[target-nums[i]]
		if ok {
			return []int{value, i}
		}

		numMap[nums[i]] = i
	}

	return []int{}
}
