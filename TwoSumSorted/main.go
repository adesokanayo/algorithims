package main

import "fmt"

func main() {

	arr := []int{2, 3, 5, 8, 11, 15}
	target := 14
	//output [1,4]
	result := TwoSumSorted(arr, target)
	fmt.Println(result)

}

func TwoSumSorted(arr []int, target int) []int {
	var final []int
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l]+arr[r] == target {
			final = append(final, l, r)
			return final
		}
		if arr[l]+arr[r] < target {
			l++
		}
		if arr[l]+arr[r] > target {
			r--
		}
	}
	return final
}
