package downloader

import (
	"music_dl/track"

	"errors"
	"os/exec"
	"strings"
)

type YoutubeDLCompatibleDownloader struct {
	Executable      string
	Format          string
	FormatExtension string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(track track.Track) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, track.URL)

	if err := dlCmd.Run(); err != nil {
		return errors.New("Download failed")
	}

	return nil
}

// replaceExtension replaces the extension of path with newExtension. 
func replaceExtension(path, newExtension string) string {
	tokens := strings.Split(path, ".")
	// Replace the final token with the new extension
	tokens[len(tokens)-1] = newExtension
	return strings.Join(tokens, ".")
}


func (ytdl YoutubeDLCompatibleDownloader) GetFilename(track track.Track) (string, error) {
	fileNameCmd := exec.Command(ytdl.Executable, "--get-filename", track.URL)
	output, err := fileNameCmd.Output()

	filename := string(output[:])

	if err != nil {
		return filename, errors.New("Could not get filename")
	}

	filename = replaceExtension(filename, ytdl.FormatExtension)

	return filename, nil
}

