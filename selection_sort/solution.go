package main

func solution(arr []int) {
	total := len(arr)
	for i := 0; i < total-1; i++ {
		minIndex := i
		for j := i; j < total; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
}
