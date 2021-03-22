package main

import "fmt"

func main() {
	unsorted := []int{2, 3, 434, 5, 45, 65, 6, 67, 67, 78, 78, 43, 462, 323}
	fmt.Println(unsorted)
	sorted := bubbleSort(unsorted)
	fmt.Println(sorted)
}

func bubbleSort(unsorted []int) []int {

	for i := 0; i < len(unsorted); i++ {
		for j := 0; j < len(unsorted)-1-i; j++ {
			if unsorted[j] > unsorted[j+1] {
				unsorted[j], unsorted[j+1] = unsorted[j+1], unsorted[j]
			}
		}
	}

	return unsorted
}

// Output
// $ go run bubble_sort.go
// [2 3 5 6 43 45 65 67 67 78 78 323 434 462]