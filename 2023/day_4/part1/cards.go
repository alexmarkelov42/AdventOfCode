package cards

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

func ConvertToInt(winNumbersStr []string) [][]int {
	arr := make([][]int, len(winNumbersStr))
	for i, str := range winNumbersStr {
		var(
			ind = 0
			tmp = 0
			err error
		)
		for ind != len(str)-1 {
			ind, err = findNextDigit(str, ind)
			if err != nil {
				break
			}
			
			_, err := fmt.Sscan(str[ind:], &tmp)
			if err == nil {
				arr[i] = append(arr[i], tmp)
			}
			ind, err = findNextNonDigit(str, ind)
		}
	}
	return arr
}

func findNextNonDigit(str string, ind int) (int, error) {
	for unicode.IsDigit(rune(str[ind])) {
		if ind == len(str)-1 {
			return ind, errors.New("index out of bounds")
		}
		ind++
	}
	return ind, nil
}

func findNextDigit(str string, ind int) (int, error) {
	for !unicode.IsDigit(rune(str[ind])) {
		if ind == len(str)-1 {
			return ind, errors.New("index out of bounds")
		}
		ind++
	}
	return ind, nil
}

func GetTotalPoints(winNumbers, cardNumbers [][]int) int {
	var (
		pointsSum = 0
	)
	for i := 0; i < len(winNumbers); i++ {
		numOfMatches := GetNumberOfMatches(winNumbers[i], cardNumbers[i])
		points := GetCardPoints(numOfMatches)
		pointsSum += points
	}
	return pointsSum
}

func GetNumberOfMatches(winNumbers, cardNumbers []int) int {
	var matches = 0
	for _, num := range cardNumbers {
		if slices.Contains(winNumbers, num) {
			matches++
		}
	}	
	return matches
}

func GetCardPoints(matches int) int {
	return int(math.Pow(2, float64(matches-1)))
}

func GetDataFromFile(filepath string) ([][]int, [][]int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var (
		winNumbers  []string
		cardNumbers []string
	)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for i := 0; scanner.Scan(); i++ {
		_, numbers, _ := strings.Cut(scanner.Text(), ": ")
		nums := strings.Split(numbers, "|")
		winNumbers = append(winNumbers, nums[0])
		cardNumbers = append(cardNumbers, nums[1])
	}
	return ConvertToInt(winNumbers), ConvertToInt(cardNumbers)
}

func GetFilePath(inputName string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(filepath.Dir(wd)) + "/" + inputName
}
