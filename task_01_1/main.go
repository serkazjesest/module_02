package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "104"
	var d int = 35

	res_s, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T \n", res_s)

	res_d := strconv.Itoa(d)

	fmt.Printf("%T \n", res_d)
}
