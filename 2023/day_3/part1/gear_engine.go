package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var N = 140

func main() {
	file, err := os.Open("../Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data := makeArrayFromFile(file)
	result := findSumOfAllParts(data)

	fmt.Println(result)
}

func findSumOfAllParts(data [][]byte) int {
	var sum = 0
	for i_str, line := range data {
		numbersIndex := getNumbersIndex(line)
		for _, numIndexes := range numbersIndex {
			result := checkNumber(i_str, numIndexes, data)
			if result {
				num := getNumberFromIndexes(line, numIndexes)
				sum += num
			}
		}
	}
	return sum
}

func getNumberFromIndexes(line []byte, numIndexes []int) int {
	i1 := numIndexes[0]
	i2 := numIndexes[1]
	slice := line[i1:i2]
	num, _ := strconv.Atoi(string(slice))
	return num
}

func getNumbersIndex(line []byte) [][]int {
	regexpNumber, _ := regexp.Compile("[0-9]+")
	numbersIndex := regexpNumber.FindAllIndex(line, -1)
	return numbersIndex
}

func checkNumber(i int, numIndexes []int, data [][]byte) bool {
	var result bool = true
	result = checkNeighbors(i, numIndexes, data)
	return result
}

func checkNeighbors(i int, numIndexes []int, data [][]byte) bool {
	for ind := numIndexes[0]; ind < numIndexes[1]; ind++ {
		indexes := []struct {
			i int
			j int
		}{
			{i - 1, ind - 1},
			{i - 1, ind},
			{i - 1, ind + 1},
			{i, ind - 1},
			{i, ind + 1},
			{i + 1, ind - 1},
			{i + 1, ind},
			{i + 1, ind + 1},
		}
		for _, kap := range indexes {
			res := checkNeighbor(kap.i, kap.j, data)
			if res == true {
				return true
			}
		}
	}

	return false
}

func checkNeighbor(i1, i2 int, data [][]byte) bool {
	if i1 < 0 || i2 < 0 || i1 > N-1 || i2 > N-1 {
		return false
	}
	if data[i1][i2] != '.' && !unicode.IsDigit(rune(data[i1][i2])) {
		return true
	}
	return false
}

func checkUp(i int, numIndexes []int, data [][]byte) bool {
	if i == 0 {
		return false
	}
	return false
}

func makeArrayFromFile(file *os.File) [][]byte {
	var data = make([][]byte, N)
	for i := 0; i < N; i++ {
		data[i] = make([]byte, N)
	}
	scanner := bufio.NewScanner(file)
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < N; j++ {
			data[i][j] = line[j]
		}
	}
	return data
}
