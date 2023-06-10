package happy_number

func isHappy(num int) bool {
	slow, fast := num, sumOfSquares(num)
	for fast != 1 {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
		if fast == slow {
			return false
		}
	}

	return true
}

func square(num int) int {
	return num * num
}

func sumOfSquares(num int) int {
	total := 0
	for num > 0 {
		digit := num % 10
		num = num / 10

		total += square(digit)
	}

	return total
}
