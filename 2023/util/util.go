package util

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func GetDefaultFilePath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(filepath.Dir(wd)) + "/" + "Input.txt"
}

func GetFilePathFromSrc(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(filepath.Dir(wd)) + "/" + filename
}

func ReadArrayFromFile(filepath string) []string {
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
