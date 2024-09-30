package day3_part1

import (
	"bufio"
	"log"
	"os"
)

type Coord struct {
	x int
	y int
}

func CountHouses(filename string) (total int) {
	houses := make([]Coord, 100)
	houses = append(houses, Coord{0, 0})
	total = 1
	var x, y int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		switch ch {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--

		}
		newHouse := Coord{x, y}
		if !ContainsHouse(newHouse, houses) {
			houses = append(houses, newHouse)
			total += 1
		}

	}
	return total
}

func ContainsHouse(coord Coord, houses []Coord) bool {
	for _, val := range houses {
		if CompareStructs(val, coord) {
			return true
		}
	}
	return false
}

func CompareStructs(val, coord Coord) bool {
	return (val.x == coord.x) && (val.y == coord.y)
}
