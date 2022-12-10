package main

import (
	"os"
	"os/exec"
)

func Download(exe string, outFmt string, url string) error {
	// TOOD: Add options which increase the chances of successful download?
	dlCmd := exec.Command(exe, "-x", "--audio-format", outFmt, url)
	return dlCmd.Run()
}

func GetDownloadedFilePath(exe string, url string) (string, error) {
	fileNameCmd := exec.Command(exe, "--get-filename", url)
	name, err := fileNameCmd.Output()
	return string(name[:]), err
}

func Move(source string, destination string) error {
	err := os.Rename(source, destination)
	return err
}
