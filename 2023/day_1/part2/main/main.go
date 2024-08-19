package main

import (
	"fmt"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_1/part2"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/util"
)

func main() {
	sum := part2.GetCalibrationSumWithWords(util.GetDefaultFilePath())
	fmt.Println(sum)
}
