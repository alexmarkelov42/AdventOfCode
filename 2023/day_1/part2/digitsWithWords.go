package part2

import (
	"bufio"
	"log"
	"os"
	"strings"

	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_1/part1"
)

var digitWords = map[string]int{
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

func GetCalibrationSumWithWords(filename string) int {
	var sum = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num1 := getFirstDigit(scanner.Text())
		num2 := getSecondDigit(scanner.Text())
		sum += num1*10 + num2
	}
	return sum
}

func getSecondDigit(line string) int {
	digit, digitIndex, _ := part1.FindSecondDigit(line)
	digitWord, digitWordIndex := FindSecondDigitWord(line)
	if digitWordIndex > digitIndex {
		return digitWord
	}
	return digit
}

func FindSecondDigitWord(line string) (int, int) {
	var (
		digitWordIndex = 0
		digitWord      = 0
	)
	for word, digit := range digitWords {
		ind := 0
		n := 0
		for ind != -1 {
			ind = strings.Index(line[n:], word)
			if ind != -1 {
				if ind+n > digitWordIndex {
					digitWordIndex = ind + n
					digitWord = digit
				}
				n = n + ind + len(word)
			}
		}
	}
	return digitWord, digitWordIndex
}

func getFirstDigit(line string) int {
	digit, digitIndex, _ := part1.FindFirstDigit(line)
	digitWord, digitWordIndex := findFirstDigitWord(line)
	if digitWordIndex < digitIndex {
		return digitWord
	}
	return digit
}

func findFirstDigitWord(line string) (int, int) {
	var (
		digitWord      = 0
		digitWordIndex = len(line)
	)
	for word, digit := range digitWords {
		ind := strings.Index(line, word)
		if ind != -1 {
			if ind < digitWordIndex {
				digitWordIndex = ind
				digitWord = digit
			}
		}
	}
	return digitWord, digitWordIndex
}
