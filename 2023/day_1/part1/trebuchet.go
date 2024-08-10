package trebuchet

import (
	"strconv"
	"unicode"
)

func findFirstDigit(documentLine string) int {
	for _, char := range documentLine {
		if unicode.IsDigit(char) {
			result, _ := strconv.Atoi(string(char))
			return result
		}
	}
	return 0 // TODO return err?
}

func findSecondDigit(documentLine string) int {
	for i := len(documentLine) - 1;  i >= 0; i-- {
		if unicode.IsDigit(rune(documentLine[i])) {
			result, _ := strconv.Atoi(string(documentLine[i]))
			return result
		}
	}
	return 0 // TODO return err?
}

func GetCalibrationSum(document []string) int {
	var sum = 0
	for _, str := range document {
		num1 := findFirstDigit(str)
		num2 := findSecondDigit(str)
		sum += num1 * 10 + num2
	}
	return sum
}
