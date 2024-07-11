package main

import (
    "downloader/downloader"
    "downloader/track"
    "os"
)

// removeExisting returns tracks that are missing from the directory.
func removeExisting(tracks []track.Track, downloader downloader.Downloader, directory string) []track.Track {
	result := make([]track.Track, 0, len(tracks))

	for _, track := range tracks {
		if !exists(downloader.GetOutputFilename(track, directory)) {
			result = append(result, track)
		}
	}

	return result
}

// exists tests whether the file found at filename is accessible.
func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// mkdir wraps os.Mkdir so that the directory is not created if it already exists.
func mkdir(directory string) error {
	if exists(directory) {
		return nil
	}

	return os.Mkdir(directory, 0777)
}
