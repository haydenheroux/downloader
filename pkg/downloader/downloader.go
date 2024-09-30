package downloader

import "github.com/haydenheroux/media/pkg/resource"

// Downloader is implemented by any type that is used to download tracks.
type Downloader interface {
	// Download performs the process of downloading the track.
	Download(resource.Resource, string) error
	// GetFilename gets the output filename for when the track is downloaded.
	GetOutputFilename(resource.Resource, string) string
}
