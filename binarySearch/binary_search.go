package main

import "fmt"

func main() {

	sorted := []int64{2, 3, 5, 6, 6, 9, 10, 67, 78, 78, 79, 323}
	fmt.Println(sorted)
	key := int64(10)
	present := binarySearch(sorted, key)
	if present {
		fmt.Println(key, "is present in the array")
	} else {
		fmt.Println(key, "not present in the array")
	}

}

func binarySearch(array []int64, key int64) bool {

	start, end := 0, len(array)-1

	for start <= end {
		middle := (start + end) / 2
		if array[middle] < key {
			start = middle + 1
		} else if array[middle] > key {
			end = middle - 1
		} else {
			return true
		}
	}
	return false

}

//output
//go run .
//[2 3 5 6 6 9 10 67 78 78 79 323]
//10 is present in the array
