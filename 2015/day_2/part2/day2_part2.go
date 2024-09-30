package day2_part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func TotalRibon(filename string) (total int) {
	var dims = make([]int, 3)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Sscanf(str, "%dx%dx%d", &dims[0], &dims[1], &dims[2])
		sort.Slice(dims, func(i, j int) bool {
			return dims[i] < dims[j]
		})
		result := 2 * dims[0] + 2 * dims[1] + dims[0] * dims[1] * dims[2]
		total += result
	}
	return total
}
