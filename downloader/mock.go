package downloader

import (
	"downloader/resource"

	"path/filepath"
)

type MockDownloader struct {
	Format string
}

func (m MockDownloader) Download(track resource.Resource, directory string) error {
	return nil
}

func (m MockDownloader) GetOutputFilename(track resource.Resource, directory string) string {
	file := track.Name() + "." + m.Format
	return filepath.Join(directory, file)
}
