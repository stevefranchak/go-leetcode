package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	if n < 2 {
		return nil
	}

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	lastSuffixProduct := 1
	for i := n - 2; i >= 0; i-- {
		lastSuffixProduct *= nums[i+1]
		result[i] = lastSuffixProduct * result[i]
	}

	return result
}
