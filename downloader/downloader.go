package downloader

import (
	"music_dl/track"
	"os"
	"path/filepath"
)

// Downloader is implemented by any type that is used to download tracks.
type Downloader interface {
	// Download performs the process of downloading the track.
	Download(track.Track) error
	// GetFilename returns the name of the file created by Download.
	GetFilename(track.Track) (string, error)
}

func DownloadTo(downloader Downloader, track track.Track, directory string) error {
	err := downloader.Download(track)
	if err != nil {
		return err
	}

	filename, err := downloader.GetFilename(track)
	if err != nil {
		return err
	}

	destination := filepath.Join(directory, track.String())
	return os.Rename(filename, destination)
}
