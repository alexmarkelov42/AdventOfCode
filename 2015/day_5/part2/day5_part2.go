package day4_part2

import (
	"bufio"
	"log"
	"os"
)

func CountStrings(filename string) (total int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if checkTwoAdjacent(data) && checkTwoSeparated(data) {
			total++
		}
	}
	return total
}

func checkTwoSeparated(data string) bool {
	for i := 0; i+2 <= len(data); i++ {
		tmp := data[i : i+2]
		for j := i + 2; j+2 <= len(data); j++ {
			str := data[j : j+2]
			if tmp == str {
				return true
			}
		}
	}
	return false
}

func checkTwoAdjacent(data string) bool {
	for i := 0; i+2 < len(data); i++ {
		tmp := data[i]
		ch := data[i+2]
		if tmp == ch {
			return true
		}
	}
	return false
}
