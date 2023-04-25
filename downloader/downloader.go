package downloader

import "music_dl/track"

// Downloader is implemented by any type that is used to download tracks.
type Downloader interface {
	// Download performs the process of downloading the track.
	Download(track.Track, string) error
	// GetFilename gets the output filename for when the track is downloaded.
	GetOutputFilename(track.Track, string) string
}
