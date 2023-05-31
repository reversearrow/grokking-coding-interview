package pair_with_target_sum

func pairWithTargetSum(input []int, target int) []int {
	left, right := 0, len(input)-1
	for left < right {
		sum := input[left] + input[right]
		switch {
		case sum == target:
			return []int{left + 1, right + 1}
		case sum < target:
			left++
		case sum > target:
			right--
		}
	}

	return []int{}
}
