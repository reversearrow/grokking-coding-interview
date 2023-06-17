package sum_numbers_in_an_array

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + nums[len(nums)-1] + sum(nums[1:len(nums)-1])
}
