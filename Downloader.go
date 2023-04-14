package main

import (
	"errors"
	"os/exec"
	"strings"
)

func changeExtension(filename string, extension string) string {
	temp := strings.Split(filename, ".")
	str := strings.Join(temp[:len(temp)-1], ".")
	return str + "." + extension
}

type Downloader interface {
	Download(TrackInfo) error
	GetFilename(TrackInfo) (string, error)
}

type YoutubeDLCompatibleDownloader struct {
	Executable      string
	Format          string
	FormatExtension string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(track TrackInfo) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, track.URL)

	if err := dlCmd.Run(); err != nil {
		return errors.New("Download failed")
	}

	return nil
}

func (ytdl YoutubeDLCompatibleDownloader) GetFilename(track TrackInfo) (string, error) {
	fileNameCmd := exec.Command(ytdl.Executable, "--get-filename", track.URL)
	output, err := fileNameCmd.Output()

	filename := string(output[:])

	if err != nil {
		return filename, errors.New("Could not get filename")
	}

	filename = changeExtension(filename, ytdl.FormatExtension)

	return filename, nil
}
