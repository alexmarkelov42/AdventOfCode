package main

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

func main() {
	filepath := getFilePath("Input.txt")
	winNumbersStr, cardNumbersStr := getDataFromFile(filepath)
	winNumbers := convertToInt(winNumbersStr)
	cardNumbers := convertToInt(cardNumbersStr)
	totalPoints := getTotalPoints(winNumbers, cardNumbers)
	fmt.Println(totalPoints)
}

func convertToInt(winNumbersStr []string) [][]int {
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

func getTotalPoints(winNumbers, cardNumbers [][]int) int {
	var (
		pointsSum = 0
	)
	for i := 0; i < len(winNumbers); i++ {
		points := getCardPoints(winNumbers[i], cardNumbers[i])
		pointsSum += points
	}
	return pointsSum
}

func getCardPoints(winNumbers, cardNumbers []int) int {
	var points = 0
	for _, num := range cardNumbers {
		if slices.Contains(winNumbers, num) {
			points++
		}
	}
	return int(math.Pow(2, float64(points-1)))
}

func getDataFromFile(filepath string) ([]string, []string) {
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
	return winNumbers, cardNumbers
}

func getFilePath(inputName string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(wd) + "/" + inputName
}
