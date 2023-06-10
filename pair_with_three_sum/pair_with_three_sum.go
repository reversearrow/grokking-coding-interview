package pair_with_three_sum

import (
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := range nums {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return true
			}

			if sum < target {
				left++
			}

			if sum > target {
				right--
			}
		}
	}

	return false
}
