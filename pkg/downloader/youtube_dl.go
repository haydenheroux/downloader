package downloader

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/haydenheroux/media/pkg/resource"
)

type YoutubeDLCompatibleDownloader struct {
	Executable string
	Format     string

	outputDirectory string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(track resource.Resource) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, "-o", track.Title(), track.Source())
	dlCmd.Dir = ytdl.outputDirectory

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

func (ytdl *YoutubeDLCompatibleDownloader) SetOutputDirectory(directory string) {
	ytdl.outputDirectory = directory
}

func (ytdl YoutubeDLCompatibleDownloader) OutputLocation(track resource.Resource) string {
	file := track.Title() + "." + ytdl.Format
	return filepath.Join(ytdl.outputDirectory, file)
}
