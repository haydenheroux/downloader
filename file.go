package main

import "os"

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
