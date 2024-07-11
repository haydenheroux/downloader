package downloader

import (
	"downloader/track"

	"path/filepath"
)

type MockDownloader struct {
	Format string
}

func (m MockDownloader) Download(track track.Track, directory string) error {
	return nil
}

func (m MockDownloader) GetOutputFilename(track track.Track, directory string) string {
	file := track.Name + "." + m.Format
	return filepath.Join(directory, file)
}
