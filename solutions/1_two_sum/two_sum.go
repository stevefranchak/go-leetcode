package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	answer := make([]int, 2)

	for i, n := range nums {
		if v, ok := m[n]; ok {
			answer = []int{v, i}
			break
		}
		m[target-n] = i
	}

	return answer
}
