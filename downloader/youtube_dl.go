package downloader

import (
	"downloader/track"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
)

type YoutubeDLCompatibleDownloader struct {
	Executable string
	Format     string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(track track.Track, directory string) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, "-o", track.Name(), track.URL())
	dlCmd.Dir = directory

	output, err := dlCmd.CombinedOutput()

	if err == nil {
		return err
	}

	return errorFromOutput(output, err)
}

func errorFromOutput(output []byte, err error) error {
	s := string(output)

	var errs []error

	if executable, ok := err.(*exec.Error); ok {
		errs = append(errs, missingDependencyError(executable.Name))
	}

	if strings.Contains(s, "ffmpeg not found") {
		errs = append(errs, missingDependencyError("ffmpeg"))
	}

	if strings.Contains(s, "Video unavailable") || strings.Contains(s, "not a valid URL") || strings.Contains(s, "looks truncated") {
		errs = append(errs, unavailableError())
	}

	if strings.Contains(s, "confirm your age") {
		errs = append(errs, ageRestricted())
	}

	if len(errs) == 0 {
		errs = append(errs, errors.New(s))
	}

	return errors.Join(errs...)
}

func (ytdl YoutubeDLCompatibleDownloader) GetOutputFilename(track track.Track, directory string) string {
	file := track.Name() + "." + ytdl.Format
	return filepath.Join(directory, file)
}
