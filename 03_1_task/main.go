package main

import "fmt"

func main() {
	days := []string{"mon", "tues", "wedn", "thur", "fri", "sat", "sun"}
	workDay := days[0:5]
	days = days[5:7]

	fmt.Println(workDay) //[mon tues wedn thur fri]
	fmt.Println(days)    //[sat sun]

}
