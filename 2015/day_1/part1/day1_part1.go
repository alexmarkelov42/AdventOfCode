package day1_part1

import (
	"bufio"
	"log"
	"os"
)

func CountFloor(filepath string) (floor int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if r == '(' {
			floor++
		}
		if r == ')' {
			floor--
		}
		if err != nil {
			break
		}
	}
	return floor
}
