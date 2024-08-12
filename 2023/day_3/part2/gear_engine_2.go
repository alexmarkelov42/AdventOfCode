package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var N = 140

type Coord struct {
	i int
	j int
}

type numberIndex struct {
	start Coord
	end   Coord
}

func main() {
	file, err := os.Open("../Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data := makeArrayFromFile(file)
	result := findSumOfAllGears(data)
	fmt.Println(result)
}

func findSumOfAllGears(data [][]byte) int {
	var sum = 0
	for row, line := range data {
		lineSum := 0
		for col, symbol := range line {
			if string(symbol) == "*" {
				ratio := gearRatio(row, col, data)
				sum += ratio
				lineSum += ratio
			}
		}
	}
	return sum
}

func gearRatio(row, col int, data [][]byte) int {
	var numbers []numberIndex
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if isArrayElementDigit(i, j, data) {
				jMin := findMinIndex(i, j, data)
				jMax := findMaxIndex(i, j, data)
				newNumberIndex := numberIndex{Coord{i, jMin}, Coord{i, jMax}}
				if !isNewNumberPresent(newNumberIndex, numbers) {
					numbers = append(numbers, newNumberIndex)
				}
			}
		}
	}
	if len(numbers) == 2 {
		num1 := getNumberFromCoords(numbers[0], data)
		num2 := getNumberFromCoords(numbers[1], data)
		fmt.Printf("%d) %d = %d * %d\n", row, num1*num2, num1, num2)
		return num1 * num2

	}
	return 0
}

func getNumberFromCoords(numberIndex numberIndex, data [][]byte) int {
	i := numberIndex.start.i
	j1 := numberIndex.start.j
	j2 := numberIndex.end.j
	number := data[i][j1:j2]
	num, _ := strconv.Atoi(string(number))
	return num
}

func isNewNumberPresent(newNumberIndex numberIndex, numbers []numberIndex) bool {
	for _, num := range numbers {
		if equalCoords(newNumberIndex.start, num.start) && equalCoords(newNumberIndex.end, num.end) {
			return true
		}
	}
	return false
}

func getNumberFromIndex(digitCoord Coord, data [][]byte) int {
	i := digitCoord.i
	j1 := findMinIndex(digitCoord.i, digitCoord.j, data)
	j2 := findMaxIndex(digitCoord.i, digitCoord.j, data)
	number := data[i][j1:j2]
	num, _ := strconv.Atoi(string(number))
	return num

}

func findMaxIndex(i, j int, data [][]byte) int {
	for {
		if j == N-1 {
			break
		}
		if unicode.IsDigit(rune(data[i][j+1])) {
			j++
		} else {
			break
		}
	}
	j++
	return j
}

func findMinIndex(i, j int, data [][]byte) int {
	for {
		if j == 0 {
			break
		}
		if unicode.IsDigit(rune(data[i][j-1])) {
			j--
		} else {
			break
		}
	}
	return j
}

func equalCoords(coord1, coord2 Coord) bool {
	return coord1.i == coord2.i && coord1.j == coord2.j
}

func getNumberFromIndexes(line []byte, numIndexes []int) int {
	i1 := numIndexes[0]
	i2 := numIndexes[1]
	slice := line[i1:i2]
	num, _ := strconv.Atoi(string(slice))
	return num
}

func isArrayElementSymbol(i1, i2 int, data [][]byte) bool {
	if i1 < 0 || i2 < 0 || i1 > N-1 || i2 > N-1 {
		return false
	}
	if data[i1][i2] == '*' {
		return true
	}
	return false
}

func isArrayElementDigit(i1, i2 int, data [][]byte) bool {
	if i1 < 0 || i2 < 0 || i1 > N-1 || i2 > N-1 {
		return false
	}
	if unicode.IsDigit(rune(data[i1][i2])) {
		return true
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
