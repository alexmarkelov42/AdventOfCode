package main

import (
	"fmt"

	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_4/part1"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_4/part2"
)

func main() {
	filepath := cards.GetFilePath("Input.txt")
	winNumbers, cardNumbers := cards.GetDataFromFile(filepath)
	totalPoints := cardWithCopies.GetTotalPointsWithCopies(winNumbers, cardNumbers)
	fmt.Println(totalPoints)
}
