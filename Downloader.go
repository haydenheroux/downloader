package main

import (
	"errors"
	"os/exec"
	"strings"
)

// replaceExtension replaces the extension of path with newExtension. 
func replaceExtension(path, newExtension string) string {
	tokens := strings.Split(path, ".")
	// Replace the final token with the new extension
	tokens[len(tokens)-1] = newExtension
	return strings.Join(tokens, ".")
}

// Downloader is implemented by any type that is used to download tracks.
type Downloader interface {
	// Download performs the process of downloading the track.
	Download(TrackInfo) error
	// GetFilename returns the name of the file created by Download.
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

	filename = replaceExtension(filename, ytdl.FormatExtension)

	return filename, nil
}
