package downloader

import (
	"music_dl/track"

	"path/filepath"
)

type MockDownloader struct {
	Format string
}

func (m MockDownloader) Download(track track.Track, directory string) error {
	return nil
}

func (m MockDownloader) GetOutputFilename(track track.Track, directory string) string {
	file := track.String() + "." + m.Format
	return filepath.Join(directory, file)
}
