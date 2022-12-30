package main

import (
	"errors"
	"os"
	"os/exec"
)

func Download(exe string, outFmt string, url string) error {
	// TOOD: Add options which increase the chances of successful download?
	dlCmd := exec.Command(exe, "-x", "--audio-format", outFmt, url)

	if err := dlCmd.Run(); err != nil {
		return errors.New("Download failed")
	} 

	return nil
}

func DownloadedFilename(exe string, url string) (string, error) {
	fileNameCmd := exec.Command(exe, "--get-filename", url)
	output, err := fileNameCmd.Output()

	filename := string(output[:])

	if err != nil {
		return filename, errors.New("Could not get filename")
	}

	return filename, err
}

func Move(source string, destination string) error {
	err := os.Rename(source, destination)
	return err
}
