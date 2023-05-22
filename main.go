package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()

	out, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "empty field":
				out, err := os.Open("out.txt")
				if err != nil {
					fmt.Println(err)
					return
				}
				defer out.Close()

				s := bufio.NewScanner(out)
				count := 0

				for s.Scan() {
					count++
					fmt.Println(s.Text())
				}
				fmt.Println(fmt.Errorf("parse error: empty field on string %d", count))

			default:
				panic("critical")
			}
		}
	}()

	s := bufio.NewScanner(in)

	line := 0
	isEmpty := false
	for s.Scan() {
		line++
		workStr := s.Text()
		workSlice := strings.SplitN(workStr, "|", 3)
		for _, s := range workSlice {
			if s == "" {
				isEmpty = true
			}
		}
		if !isEmpty {
			workStr = fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", line, workSlice[0], workSlice[1], workSlice[2])
			out.WriteString(workStr)

		} else {
			panic("empty field")
		}
	}
}
