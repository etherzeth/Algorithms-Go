package main

func twoSum(nums []int, target int) []int {
	// nums [1, 2] & target 3. in Index 1 = 0, 2 = 1
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := mymap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		mymap[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}
