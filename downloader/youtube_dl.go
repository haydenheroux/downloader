package downloader

import (
	"music_dl/track"

	"os/exec"
	"path/filepath"
)

type YoutubeDLCompatibleDownloader struct {
	Executable      string
	Format          string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(track track.Track, directory string) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, "-o", track.String(), track.URL)
	dlCmd.Dir = directory

	return dlCmd.Run()
}

func (ytdl YoutubeDLCompatibleDownloader) GetOutputFilename(track track.Track, directory string) string {
	file := track.String() + "." + ytdl.Format
	return filepath.Join(directory, file)
}
