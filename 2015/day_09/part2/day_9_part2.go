package day_9_part1

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const N int = 8

var shortestPath int

func LongestRoute(filename string) int {
	data := readArray(filename)
	longestPath := findLongestPath(data)
	return longestPath
}

func findLongestPath(data [][]int) (total int) {
	arr := make([]int, N)
	for i := range N {
		arr[i] = i
	}
	total = allPermutations(arr, data, N, N, total)
	return total
}

func allPermutations(arr []int, data [][]int, size, n, total int) (result int) {
	if size == 1 {
		res := pathLength(arr, data)
		if res > total || total == 0 {
			total = res
		}
		return total
	}
	for i := range(N) {
		total = allPermutations(arr, data, size - 1, n, total)
		if size % 2 == 1 {
			arr[0], arr[size-1] = arr[size-1], arr[0]
		} else {
			arr[i], arr[size-1] = arr[size-1], arr[i]
		}
	}
	return total
}

func pathLength(arr []int, data [][]int) (result int) {
	for i := 1; i < len(arr); i++ {
		i1 := arr[i-1]
		i2 := arr[i]
		result += data[i1][i2]
	}
	return result
}

func readArray(filename string) (arr [][]int) {
	data := make([][]int, N)
	for i := range N {
		data[i] = make([]int, N)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := range N {
		for j := i + 1; j < N; j++ {
			scanner.Scan()
			str := strings.Split(scanner.Text(), " = ")
			data[i][j], _ = strconv.Atoi(str[1])
			data[j][i] = data[i][j]
		}
	}
	return data
}

func findMinInd(arr []int, path map[int]bool) (minEl, ind int) {
	minEl = slices.Max(arr)
	for i := 0; i < len(arr); i++ {
		if !path[i] {
			if arr[i] != 0 && arr[i] <= minEl {
				minEl = arr[i]
				ind = i
			}
		}
	}
	return minEl, ind
}
