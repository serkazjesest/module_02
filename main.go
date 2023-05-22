package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var startTime = time.Now()

func logTime() {
	fmt.Println(time.Since(startTime))
}

func main() {
	defer logTime()
	in, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Эту обработку ошибок меня заставила делать Кашкина А.О. Файл короче не открывается.")
		return
	}
	defer in.Close()

	out, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Эту обработку ошибок меня заставила делать Кашкина А.О. Файл короче не создаётся.")
		return
	}
	defer out.Close()

	s := bufio.NewScanner(in)
	count := 0
	bytes := 0
	for s.Scan() {
		count++
		strForWrite := strconv.Itoa(count) + ". " + s.Text() + "\n"
		bytes = bytes + len(strForWrite)
		out.WriteString(strForWrite)
	}
	defer fmt.Print("Количество строк записанных в файл:", count, "\n", "Количество байт записанных в файл:", bytes, "\n")
}
