package day5_part1

import (
	"bufio"
	"log"
	"os"
	"strings"
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
		if checkIfLegal(data) && checkVowels(data) && checkDouble(data) {
			total++
		}

	}
	return total
}

func checkDouble(data string) bool {
	for i := 1; i < len(data); i++ {
		if data[i] == data[i-1] {
			return true
		}
	}
	return false
}

func checkVowels(data string) bool {
	count := 0
	for i := 0; i < len(data); i++ {
		if strings.Contains("aeoui", string(data[i])) {
			count++
		}
		if count >= 3 {
			return true
		}
	}
	return false
}

func checkIfLegal(data string) bool {
	if strings.Contains(data, "ab") || strings.Contains(data, "cd") || strings.Contains(data, "pq") || strings.Contains(data, "xy") {
		return false
	}
	return true
}
