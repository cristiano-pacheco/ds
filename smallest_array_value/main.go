package main

// findSmallest returns a pointer to the index of the smallest element in the given integer slice.
// If the slice is empty, it returns nil.
func findSmallest(arr []int) *int {
	if len(arr) == 0 {
		return nil
	}

	smallest := arr[0]
	smallest_index := 0
	total := len(arr)

	for i := 1; i < total; i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return &smallest_index
}
