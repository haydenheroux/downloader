package downloader

import (
	"path/filepath"

	"github.com/haydenheroux/media/pkg/resource"
)

type MockDownloader struct {
	Format string
}

func (m MockDownloader) Download(track resource.Resource, directory string) error {
	return nil
}

func (m MockDownloader) GetOutputFilename(track resource.Resource, directory string) string {
	file := track.Title() + "." + m.Format
	return filepath.Join(directory, file)
}
