package main

func removeDuplicates(nums []int) int {
	length := 0
	if len(nums) == length {
		return length
	}
	length = 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[length] = nums[i]
			length++
		}
	}
	nums = nums[:length]
	return length
}
