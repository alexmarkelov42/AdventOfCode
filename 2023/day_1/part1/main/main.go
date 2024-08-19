package main

import (
	"fmt"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_1/part1"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/util"
)

func main() {
	result := part1.GetCalibrationSum(util.GetDefaultFilePath())
	fmt.Println(result)
}
