package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cubes := getSetsOfCubes(line)
		result += findPowerOfCubeSet(cubes)
	}
	fmt.Println(result)
}

func findPowerOfCubeSet(cubes [][]byte) int {
	mins := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, token := range cubes {
		number := getNumber(token)
		color := getColor(token)
		if mins[color] == 0 {
			mins[color] = number
		}
		if number > mins[color] {
			mins[color] = number
		}
	}
	return mins["red"] * mins["blue"] * mins["green"]
}

func getGameId(input string) int {
	regexpGameId, _ := regexp.Compile("[0-9]+")
	GameByte := regexpGameId.Find([]byte(input))
	GameId, _ := strconv.Atoi(string(GameByte))
	return GameId
}

func getSetsOfCubes(line string) [][]byte {
	regexpListOfCubes, _ := regexp.Compile("([0-9]+ (blue|green|red)){1}")
	listOfCubes := regexpListOfCubes.FindAll([]byte(line), -1)
	return listOfCubes
}

func getNumber(token []byte) int {
	regexprNumber, _ := regexp.Compile("[0-9]+")
	numberByte := regexprNumber.Find([]byte(token))
	number, _ := strconv.Atoi(string(numberByte))
	return number
}

func getColor(token []byte) string {
	regexprColor, _ := regexp.Compile("(blue|green|red)")
	color := regexprColor.Find([]byte(token))
	return string(color)
}

func isResultValid(number int, color string) bool {
	result := []struct {
		amount int
		color  string
	}{
		{12, "red"},
		{13, "green"},
		{14, "blue"},
	}
	for _, res := range result {
		if color == res.color {
			if number > res.amount {
				return false
			}
		}
	}
	return true
}
