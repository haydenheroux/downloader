package downloader

import (
	"path/filepath"

	"github.com/haydenheroux/media/pkg/resource"
)

type MockDownloader struct {
	Format string

	outputDirectory string
}

func (m MockDownloader) Download(track resource.Resource) error {
	return nil
}

func (m *MockDownloader) SetOutputDirectory(directory string) {
	m.outputDirectory = directory
}

func (m MockDownloader) OutputLocation(track resource.Resource) string {
	file := track.Title() + "." + m.Format
	return filepath.Join(m.outputDirectory, file)
}
