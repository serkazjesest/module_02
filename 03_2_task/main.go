package main

import "fmt"

func main() {
	workDays := []string{"mon", "tues", "wedn", "thur", "fri"}
	weekEndDays := []string{"sat", "sun"}

	days := make([]string, 0, 7)
	days = append(days, workDays...)
	days = append(days, weekEndDays...)

	fmt.Println(days) //[mon tues wedn thur fri sat sun]
}
