package main

func main() {

}

func removeDup(a []string, b []string) (c []string) {

	slice := make([]int, 0)
	checker := make(map[int]int)

	d := append(a, b...)
	result := make([]int, 0)

	for _, val := range d {
		slice[val] = 1
	}

	for ful := range slice{
		append(result, ful)
	}
	return result
}
