package squares_of_sorted_array

func square(i int) int {
	return i * i
}

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1

	var output = make([]int, len(nums), len(nums))
	counter := len(nums) - 1
	for left <= right {
		leftSqure := square(nums[left])
		rightSqure := square(nums[right])

		switch {
		case leftSqure > rightSqure:
			output[counter] = leftSqure
			left++
		default:
			output[counter] = rightSqure
			right--
		}
		counter--
	}

	return output
}
