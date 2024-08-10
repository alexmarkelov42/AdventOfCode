package util

import (
	"bufio"
	"log"
	"os"
)

func GetArrayFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var document []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		document = append(document, line)
	}
	return document
}
