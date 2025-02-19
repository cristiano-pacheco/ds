package main

func sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	current := numbers[0]

	return current + sum(numbers[1:])
}

func maximum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	maxRest := maximum(numbers[1:])

	if numbers[0] > maxRest {
		return numbers[0]
	}

	return maxRest
}
