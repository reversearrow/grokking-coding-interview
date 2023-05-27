package remove_duplicates

func remove(nums []int, index int) []int {
	return append(nums[:index], nums[index+1:]...)
}

func removeDuplicates(nums []int) []int {
	start := 0
	next := start + 1

	switch len(nums) {
	case 0:
		return nums
	case 1:
		return nums
	}

	for next <= len(nums)-1 {
		if nums[start] != nums[next] {
			start = next
			next = start + 1
			continue
		}

		nums = remove(nums, next)
	}

	return nums
}
