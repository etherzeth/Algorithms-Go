func twoSum(nums []int, target int) []int {
	// Brute force

	// Time: O(n²)
	// Space: O(1)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}