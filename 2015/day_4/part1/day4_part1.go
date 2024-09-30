package day4_part1

import (
	"crypto/md5"
	"strconv"
)

func FindHash(secretKey string) string {
	var i = 0
	for {
		newString := secretKey + strconv.Itoa(i)
		hash := md5.Sum([]byte(newString))
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 15 {
			return newString
		} else {
			i++
		}
	}
}
