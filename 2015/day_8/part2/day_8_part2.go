package day_8_part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func CountChars(filename string) (total int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		qstr := strconv.Quote(data)
		total += len(qstr) - len(data)
	}
	return total
}
