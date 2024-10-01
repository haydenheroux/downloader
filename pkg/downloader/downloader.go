package downloader

import "github.com/haydenheroux/media/pkg/resource"

type Downloader interface {
	// Download downloads the resource.
	Download(resource.Resource) error
	// SetOutputDirectory sets the directory where the downloader outputs resources.
	SetOutputDirectory(string)
	// OutputLocation returns where the downloader outputs the resource.
	OutputLocation(resource.Resource) string
}
