package day6_part1

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func CountLights(filename string) (total int) {
	var N int = 1000
	var lights = make([][]bool, N)
	for i := range N {
		lights[i] = make([]bool, N)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		action := parseAction(data)
		numbers := parseNumbers(data)
		for i := numbers[0]; i <= numbers[2] || i == numbers[0]; i++ {
			for j := numbers[1]; j <= numbers[3] || j == numbers[1]; j++ {
				lights[i][j] = updateValue(action, lights[i][j])
			}
		}
	}
	for i := range N {
		for j := range N {
			if lights[i][j] == true {
				total += 1
			}
		}
	}
	return total
}

func updateValue(action string, val bool) (result bool) {
	switch action {
	case "turn on":
		result = true
	case "turn off":
		result = false
	case "toggle":
		result = !val
	}
	return result
}

func parseAction(data string) string {
	regexpAction, _ := regexp.Compile("(toggle|turn on|turn off)")
	action := regexpAction.Find([]byte(data))
	result := string(action)
	return result
}

func parseNumbers(data string) []int {
	regexpNumbers, _ := regexp.Compile("[0-9]+")
	list := regexpNumbers.FindAll([]byte(data), -1)
	numbers := make([]int, 4)
	for i, val := range list {
		tmp, _ := strconv.Atoi(string(val))
		numbers[i] = tmp
	}
	return numbers
}
