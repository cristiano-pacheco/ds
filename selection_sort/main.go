package main

import "fmt"

func main() {
	input := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Unsorted: %+v", input)
	selectionSort(input)
	fmt.Printf("sorted: %+v", input)
}

func selectionSort(arr []int) {
	// change me
}
