package main

import (
	"music_dl/downloader"
	"music_dl/track"

	"os"
	"path/filepath"
)

func downloadTo(downloader downloader.Downloader, track track.Track, directory string) error {
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
