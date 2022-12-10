package main

import (
	"bufio"
	"os"
)

func Lines(filePath string) []string {
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

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func CreateDir(dirPath string) error {
	if Exists(dirPath) {
		return nil
	}

	return os.Mkdir(dirPath, 0777)
}
