package day_8_part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func CountChars(filename string) (total int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		qstr, _ := strconv.Unquote(data)
		total = total + len(data) - len(qstr)
	}
	return total
}

func countActualChars(data string) (result int) {
	var ch rune
	for i := 1; i < len(data)-1; i++ {
		ch = rune(data[i])
		if unicode.IsLetter(ch) {
			result++
			continue
		}
		if ch == '\\' && (data[i+1] == '\\' || data[i+1] == '"') {
			result++
			i++
			continue
		}
		if ch == '\\' && data[i+1] == 'x' {
			result++
			i += 3
			continue
		}
	}
	return result
}
