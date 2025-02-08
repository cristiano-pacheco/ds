package main

import "errors"

func solution(arr []int, target int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == target {
			return mid, nil
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, errors.New("not found")
}
