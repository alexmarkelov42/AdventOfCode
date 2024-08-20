package main

import (
	"fmt"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_2/part2"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/util"
)

func main() {
	result := part2.FindSumOfAllMinimumSets(util.GetDefaultFilePath())
	fmt.Println(result)
}
