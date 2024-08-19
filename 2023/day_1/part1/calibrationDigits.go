package part1

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"unicode"
)

func FindFirstDigit(documentLine string) (int, int, error) {
	for i, char := range documentLine {
		if unicode.IsDigit(char) {
			result, _ := strconv.Atoi(string(char))
			return result, i, nil
		}
	}
	return 0, len(documentLine), errors.New("No digit found")
}

func FindSecondDigit(documentLine string) (int, int, error) {
	for i := len(documentLine) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(documentLine[i])) {
			result, _ := strconv.Atoi(string(documentLine[i]))
			return result, i, nil
		}
	}
	return 0, -1, errors.New("No digit found")
}

func GetCalibrationSum(filename string) int {
	var sum = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num1, _, _ := FindFirstDigit(scanner.Text())
		num2, _, _ := FindSecondDigit(scanner.Text())
		sum += num1*10 + num2
	}
	return sum
}
