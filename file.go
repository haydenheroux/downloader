package main

import (
	"bufio"
	"os"
)

func linesIn(filename string) ([]string, error) {
	lines := []string{}

	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func mkdir(directory string) error {
	if exists(directory) {
		return nil
	}

	return os.Mkdir(directory, 0777)
}
