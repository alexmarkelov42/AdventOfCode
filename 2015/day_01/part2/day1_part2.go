package day1_part2

import (
	"bufio"
	"log"
	"os"
)

func NegativeFloorPos(filepath string) (count int) {
	var floor int = 0
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		count++
		if r == '(' {
			floor++
		}
		if r == ')' {
			floor--
		}
		if floor < 0 {
			break
		}
		if err != nil {
			break
		}
	}
	return count
}
