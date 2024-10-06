package day2_part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func TotalSquareFeet(filename string) (total int) {
	var (
		l = 0
		w = 0
		h = 0
	)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Sscanf(str, "%dx%dx%d", &l, &w, &h)
		num1 := l * w
		num2 := l * h
		num3 := w * h
		square := 2*num1 + 2*num2 + 2*num3 + min(num1, num2, num3)
		total += square
	}
	return total
}
