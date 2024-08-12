package calibrationDigits

import (
	"errors"
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

func GetCalibrationSum(document []string) int {
	var sum = 0
	for _, str := range document {
		num1, _,  _ := FindFirstDigit(str)
		num2, _, _ := FindSecondDigit(str)
		sum += num1*10 + num2
	}
	return sum
}
