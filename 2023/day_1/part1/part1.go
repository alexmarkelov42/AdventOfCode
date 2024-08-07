package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	var (
		sum  = 0
		num1 = 0
		num2 = 0
	)
	file, err := os.Open("../Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		for _, char := range runes {
			if unicode.IsDigit(char) {
				num1 = int(char - '0')
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				num2 = int(runes[i] - '0')
				break
			}
		}
		sum += num1 * 10 + num2
	}
	fmt.Println(sum)
}
