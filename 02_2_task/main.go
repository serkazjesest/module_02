package main

import (
	"fmt"
	"math"
)

func main() {
	rad := 35 / (2 * 3.14)
	var r *float64 = &rad

	s := 3.14 * ((*r) * (*r))

	fmt.Println(math.Round(s*100) / 100)

}
