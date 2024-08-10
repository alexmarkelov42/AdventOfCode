package main

import (
	"fmt"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/day_1/part1"
	"gitlab.com/alexmarkelov42/AdventOfCode/2023/util"
)

func main() {
	filepath := "../../Input.txt"
	document := util.GetArrayFromFile(filepath)
	result := trebuchet.GetCalibrationSum(document)
	fmt.Println(result)
}

