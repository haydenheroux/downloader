package main

import (
	"bufio"
	"os"
)

func GetLines(filePath string) []string {
	var lines []string

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic("os.Open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func DoesExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func CreateDir(dir string) error {
	if DoesExist(dir) == false {
		return os.Mkdir(dir, 0777)
	}
	return nil
}
