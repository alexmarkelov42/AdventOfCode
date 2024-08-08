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

type Coord struct {
	i int
	j int
}

func main() {
	file, err := os.Open("../Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data := makeArrayFromFile(file)
	result := findSumOfAllConnectedParts(data)
	fmt.Println(result)
}

func findSumOfAllConnectedParts(data [][]byte) int {
	var sum = 0
	for i_str, line := range data {
		allNumbersIndexes := getNumbersIndex(line)
		var current_line_sum = 0
		for _, numberIndexes := range allNumbersIndexes {
			
			if i_str == 134 {
				//fmt.Println("Kappa")
			}
			result, firstDigitCoords, symbolCoords := isNeighborSymbol(i_str, numberIndexes, data)
			if result {
				isConnected, secondNumberCoord := isNeighborDigit(firstDigitCoords, symbolCoords, data)

				if isConnected {
					num1 := getNumberFromIndex(firstDigitCoords, data)
					num2 := getNumberFromIndex(secondNumberCoord, data)
					ratio := num1 * num2
					current_line_sum += ratio
					sum += ratio
				} else {
					result , firstDigitCoords, secondNumberCoord := checkUpperRight(i_str, numberIndexes[1], data)
					if result {
						num1 := getNumberFromIndex(firstDigitCoords, data)
						num2 := getNumberFromIndex(secondNumberCoord, data)
						ratio := num1 * num2
						current_line_sum += ratio
						sum += ratio
					}					
				}
			} else {
				result , firstDigitCoords, secondNumberCoord := checkUpperRight(i_str, numberIndexes[1], data)
				if result {
					num1 := getNumberFromIndex(firstDigitCoords, data)
					num2 := getNumberFromIndex(secondNumberCoord, data)
					ratio := num1 * num2
					current_line_sum += ratio
					sum += ratio
				}
			}
		}
		fmt.Printf("%d current line sum is %d\n", i_str + 1, current_line_sum)
	}
	return sum
}

func checkUpperRight(i_str int, rightDigitIndex int, data [][]byte) (bool, Coord, Coord){
	digitCoords := Coord{i_str, rightDigitIndex - 1}
	res1 := isArrayElementSymbol(digitCoords.i - 1, digitCoords.j + 1, data)
	if res1 {
		res2 := isArrayElementDigit(digitCoords.i, digitCoords.j + 2, data)
		if res2 {
			return true, digitCoords, Coord{digitCoords.i, digitCoords.j + 2}
		}
	}
	res1 = isArrayElementSymbol(digitCoords.i + 1, digitCoords.j + 1, data)
	if res1 {
		res2 := isArrayElementDigit(digitCoords.i, digitCoords.j + 2, data)
		if res2 {
			return true, digitCoords, Coord{digitCoords.i, digitCoords.j + 2}
		}
	}
	return false, Coord{}, Coord{}
}

func getNumberFromIndex(digitCoord Coord, data [][]byte) int {
	i := digitCoord.i
	j1 := digitCoord.j
	j2 := digitCoord.j
	for {
		if j1 == 0 {
			break
		}
		if unicode.IsDigit(rune(data[i][j1-1])) {
			j1--
		} else {
			break
		}
	}
	for {
		if j2 == N-1 {
			break
		}
		if unicode.IsDigit(rune(data[i][j2+1])) {
			j2++
		} else {
			break
		}
	}
	j2++
	number := data[i][j1:j2]
	num, _ := strconv.Atoi(string(number))
	return num

}

func isNeighborSymbol(i_str int, numberIndexes []int, data [][]byte) (bool, Coord, Coord) {
	for i := numberIndexes[0]; i < numberIndexes[1]; i++ {
		digitCoords := Coord{i_str, i}
		checkIndexes := []Coord{
			{digitCoords.i, digitCoords.j - 1}, // left
			{digitCoords.i + 1, digitCoords.j - 1}, // left down
			{digitCoords.i + 1, digitCoords.j},	// down
			{digitCoords.i + 1, digitCoords.j + 1}, // down right
			{digitCoords.i, digitCoords.j + 1},	// right
		}
		for _, val := range checkIndexes {
			res := isArrayElementSymbol(val.i, val.j, data)
			if res {
				return true, digitCoords, val
			}
		}
	}
	return false, Coord{}, Coord{}
}

func isNeighborDigit(firstDigitCoords, symbolCoords Coord, data [][]byte) (bool, Coord) {
	checkIndexes := []Coord{
		{symbolCoords.i + 1, symbolCoords.j - 1}, // left down
		{symbolCoords.i + 1, symbolCoords.j},	  // down
		{symbolCoords.i + 1, symbolCoords.j + 1}, // down right
		{symbolCoords.i, symbolCoords.j + 1},	  // right
	}
	for _, val := range checkIndexes {
		res := isArrayElementDigit(val.i, val.j, data)
		if res && !equalCoords(firstDigitCoords, Coord{val.i, val.j}) {
			return true, Coord{val.i, val.j}
		}
	}
	return false, Coord{}
}

func checkNumberConnectsToSymbol(digitCoords Coord, data [][]byte) (bool, Coord) {
	checkIndexes := []Coord{
		{digitCoords.i + 1, digitCoords.j - 1},
		{digitCoords.i + 1, digitCoords.j},
		{digitCoords.i + 1, digitCoords.j + 1},
		{digitCoords.i, digitCoords.j + 1},
	}
	for _, val := range checkIndexes {
		res := isArrayElementSymbol(val.i, val.j, data)
		if res {
			return true, Coord{val.i, val.j}
		}
	}
	return false, Coord{}
}

func equalCoords(coord1, coord2 Coord) bool {
	return coord1.i == coord2.i && coord1.j == coord2.j
}

func findSumOfAllParts(data [][]byte) int {
	var sum = 0
	for i_str, line := range data {
		numbersIndex := getNumbersIndex(line)
		for _, numIndexes := range numbersIndex {
			result, _ := checkNumberNeighborsSymbol(i_str, numIndexes, data)
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

func checkNumberNeighborsSymbol(i_str int, numIndexes []int, data [][]byte) (bool, Coord) {
	for ind := numIndexes[0]; ind < numIndexes[1]; ind++ {
		result, symbolCoords := checkNeighbors(i_str, ind, data)
		if result == true {
			return result, symbolCoords
		}
	}
	return false, Coord{}
}

func checkNeighbors(i_str, ind int, data [][]byte) (bool, Coord) {
	checkIndexes := []struct {
		i int
		j int
	}{
		{i_str - 1, ind - 1},
		{i_str - 1, ind},
		{i_str - 1, ind + 1},
		{i_str, ind - 1},
		{i_str, ind + 1},
		{i_str + 1, ind - 1},
		{i_str + 1, ind},
		{i_str + 1, ind + 1},
	}
	for _, kap := range checkIndexes {
		res := isArrayElementSymbol(kap.i, kap.j, data)
		if res == true {
			return true, Coord{kap.i, kap.j}
		}
	}

	return false, Coord{}
}

func isArrayElementSymbol(i1, i2 int, data [][]byte) bool {
	if i1 < 0 || i2 < 0 || i1 > N-1 || i2 > N-1 {
		return false
	}
	if data[i1][i2] != '.' && !unicode.IsDigit(rune(data[i1][i2])) {
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
