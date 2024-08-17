package main

import (
	"fmt"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_1/part2"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/util"
)

func main() {
	filepath := util.GetFilePathFromSrc("Input.txt")
	array := util.ReadArrayFromFile(filepath)
	sum := part2.GetCalibrationSumWithWords(array)
	fmt.Println(sum)
}
