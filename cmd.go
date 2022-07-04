package main

import (
	"os"
	"os/exec"
	"fmt"
)

func FindOutputFileName(exe string, url string) (string, error) {
	fileNameCmd := exec.Command(exe, "--get-filename", url)

	name, err := fileNameCmd.Output()
	if err != nil {
		msg := fmt.Sprintf("FindOutputFileName: %s", url)
		printError(msg, err)
	}
	return string(name[:]), err
}

func Move(source string, destination string) error {
	err := os.Rename(source, destination)
	if err != nil {
		msg := fmt.Sprintf("Move: %s", destination)
		printError(msg, err)
	}
	return err
}

func Download(exe string, outFmt string, url string) error {
	// Maybe add options which increase the chances of successful download?
	dlCmd := exec.Command(exe, "-x", "--audio-format", outFmt, url)

	err := dlCmd.Run()
	if err != nil {
		msg := fmt.Sprintf("Download: %s", url)
		printError(msg, err)
	}
	return err
}
