package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64

	var v1 AmericanVelocity = 120.4 * 3.6
	var v2 EuropeanVelocity = 130 * 2.2374

	convV2 := math.Round(100*float64(v2)) / 100
	v2 = EuropeanVelocity(convV2)

	fmt.Println(v1)
	fmt.Println(v2)
}
