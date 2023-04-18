package main

import (
	"bufio"
	"os"
)

// filesIn reads the all lines in the file.
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

// exists tests whether the file found at filename is accessible.
func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// mkdir wraps os.Mkdir so that the directory is not created if it already exists.
func mkdir(directory string) error {
	if exists(directory) {
		return nil
	}

	return os.Mkdir(directory, 0777)
}
