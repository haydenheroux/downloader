package downloader

import (
	"music_dl/track"

	"os/exec"
)

type YoutubeDLCompatibleDownloader struct {
	Executable      string
	Format          string
	OutputDirectory string
}

func (ytdl YoutubeDLCompatibleDownloader) Download(track track.Track) error {
	dlCmd := exec.Command(ytdl.Executable, "-x", "--audio-format", ytdl.Format, "-o", track.String(), track.URL)
	dlCmd.Dir = ytdl.OutputDirectory

	return dlCmd.Run()
}

func (ytdl YoutubeDLCompatibleDownloader) GetOutputFilename(track track.Track) string {
	return track.String() + "." + ytdl.Format
}
