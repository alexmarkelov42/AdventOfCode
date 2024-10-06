package day_10_part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LookAndSay(filename string, times int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Text()
	result := data
	for i := 0; i < times; i++ {
		result = makeSequence(result)
	}
	return len(result)
}

func makeSequence(data string) string {
	var b strings.Builder
	for i := 0; i < len(data); {
		ind := 1
		j := i
		for ; j < len(data)-1 && data[j] == data[j+1]; j++ {
			ind++
		}
		//newStr := strconv.Itoa(ind) + data[i:i+1]
		b.Write([]byte(strconv.Itoa(ind)))
		b.Write([]byte(data[i:i+1]))
		i += ind
	}
	return b.String()
}
