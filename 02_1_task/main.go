package main

import "fmt"

func main() {
	var a *int
	var b int
	a = &b
	fmt.Println(*a)
	*a = 130
	fmt.Println(b)
}
