package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum += nums[i]
		if currentSum < nums[i] {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
