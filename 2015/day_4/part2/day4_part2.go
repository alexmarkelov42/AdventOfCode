package day4_part2

import (
	"bufio"
	"crypto/md5"
	"log"
	"os"
	"strconv"
)

func FindHash(filename string) int {
	//I am reading the secret key from file because the AoC author asked not to share input (¯\_(ツ)_/¯)
	file, err := os.Open("../../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	secretKey := scanner.Text()
	var i = 0
	for {
		newString := secretKey + strconv.Itoa(i)
		hash := md5.Sum([]byte(newString))
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return i
		} else {
			i++
		}
	}
}
