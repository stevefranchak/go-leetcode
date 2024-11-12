package contains_duplicate

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}
