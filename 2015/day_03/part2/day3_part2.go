package day3_part2

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
	var x1, x2, y1, y2 int
	i := 1
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
			if i%2 == 0 {
				x2++
			} else {
				x1++
			}
		case '<':
			if i%2 == 0 {
				x2--
			} else {
				x1--
			}
		case '^':
			if i%2 == 0 {
				y2++
			} else {
				y1++
			}
		case 'v':
			if i%2 == 0 {
				y2--
			} else {
				y1--
			}

		}
		var newHouse Coord
		if i%2 == 0 {
			newHouse = Coord{x2, y2}
		} else {
			newHouse = Coord{x1, y1}
		}
		if !ContainsHouse(newHouse, houses) {
			houses = append(houses, newHouse)
			total += 1
		}
		i++
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
