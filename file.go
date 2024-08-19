package main

import (
	"os"
)

// exists tests if a file is accessible.
func exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// mkdir creates a directory if it does not already exist.
func mkdir(name string) error {
	if exists(name) {
		return nil
	}

	return os.Mkdir(name, 0777)
}
