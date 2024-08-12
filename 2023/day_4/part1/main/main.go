package main

import (
	"fmt"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_4/part1"
)

func main() {
	filepath := cards.GetFilePath("Input.txt")
	winNumbers, cardNumbers := cards.GetDataFromFile(filepath)
	totalPoints := cards.GetTotalPoints(winNumbers, cardNumbers)
	fmt.Println(totalPoints)
}
