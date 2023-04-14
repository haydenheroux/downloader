package main

import (
	"errors"
	"os/exec"
)

type Downloader interface {
	Download(TrackInfo) error
	GetFilename(TrackInfo) (string, error)
}

type YoutubeDLCompatibleDownloader struct {
	Executable string
	Format string
	FormatExtension string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(t TrackInfo) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, t.URL)

	if err := dlCmd.Run(); err != nil {
		return errors.New("Download failed")
	} 

	return nil
}

func (ytdl YoutubeDLCompatibleDownloader) GetFilename(t TrackInfo) (string, error) {
	fileNameCmd := exec.Command(ytdl.Executable, "--get-filename", t.URL)
	output, err := fileNameCmd.Output()

	filename := string(output[:])

	if err != nil {
		return filename, errors.New("Could not get filename")
	}

	filename = ChangeExtension(filename, ytdl.FormatExtension)

	return filename, nil
}
