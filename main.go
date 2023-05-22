package main

import "fmt"

func contains(a []string, x string) bool {
	for _, s := range a {
		if s == x {
			return true
		}
	}
	return false
}

func getMax(a ...int) int {
	max := a[0]
	for _, dig := range a {
		if max < dig {
			max = dig
		}
	}
	return max
}

func main() {
	fmt.Println(contains([]string{"asd", "qwe", "qwe", ""}, "q"))   // false
	fmt.Println(contains([]string{"asd", "qwe", "qwe", ""}, "qwe")) // true

	fmt.Println(getMax(1, 2, 3, -2, 0, 1241)) //1241
}
