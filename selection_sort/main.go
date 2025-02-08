package main

import "fmt"

func main() {
	input := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Unsorted: %+v", input)
	selectionSort(input)
	fmt.Printf("sorted: %+v", input)
}

func selectionSort(arr []int) {
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
