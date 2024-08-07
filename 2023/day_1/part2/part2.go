package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var (
		sum = 0
	)
	file, err := os.Open("../Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num1 := searchFirstDigit(line)
		num2 := searchSecondDigit(line)
		sum += num1*10 + num2
	}
	fmt.Println(sum)
}

func searchFirstDigit(line string) int {
	num1 := 0
	digit_index := len(line)
	word_index := len(line)
	digit_found := 0
	runes := []rune(line)
	for i, char := range runes {
		if unicode.IsDigit(char) {
			num1 = int(char - '0')
			digit_index = i
			break
		}

	}
	for word, digit := range digits {
		ind := strings.Index(line, word)
		if ind != -1 {
			if ind < word_index {
				word_index = ind
				digit_found = digit
			}
		}
	}
	if word_index < digit_index {
		num1 = digit_found
	}
	return num1
}

func searchSecondDigit(line string) int {
	num2 := 0
	digit_index := 0
	word_index := 0
	digit_found := 0
	runes := []rune(line)
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			num2 = int(runes[i] - '0')
			digit_index = i
			break
		}
	}
	for word, digit := range digits {
		ind := 0
		tmp := line
		n := 0		
		for ind != -1 {
			ind = strings.Index(tmp[n:], word)
			if ind != -1 {
				if ind + n > word_index {
					word_index = ind + n
					digit_found = digit
				}
				n = n + ind + len(word)
			}
		}
	}
	if word_index > digit_index {
		num2 = digit_found
	}
	return num2
}
