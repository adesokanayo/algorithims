package main

import (
	"fmt"
)

func main() {

	var arr1 []int
	var arr2 []string
	arr1 = append(arr1, 0, 1, 2, 3, 4, 5, 5, 6, 6, 7)
	arr2 = append(arr2, "a", "b", "b", "c", "d", "e", "f")
	fmt.Println(removeDupInt(arr1))

	fmt.Println(removeDupString(arr2))

}

// removeDupInt is based on a two pointer method.
func removeDupInt(arr []int) (result []int) {

	slow := 0
	for fast := 0; fast < len(arr); fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}
	return arr[:slow+1]
}

func removeDupString(arr []string) (result []string) {

	slow := 0
	for fast := 0; fast < len(arr)-1; fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}
	return arr[0 : slow+1]
}
