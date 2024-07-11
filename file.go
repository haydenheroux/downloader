package main

import (
	"downloader/downloader"
	"downloader/track"
	"os"
)

// scanExisting computes the set of tracks existing in a directory.
func scanExisting(tracks []track.Track, downloader downloader.Downloader, directory string) map[track.Track]bool {
	existing := make(map[track.Track]bool)

	for _, track := range tracks {
		if exists(downloader.GetOutputFilename(track, directory)) {
			existing[track] = true
		}
	}

	return existing
}

// removeExisting removes existing tracks from a slice of tracks.
func removeExisting(tracks []track.Track, existing map[track.Track]bool) []track.Track {
	result := make([]track.Track, 0, len(tracks))

	for _, track := range tracks {
		if exists, _ := existing[track]; !exists {
			result = append(result, track)
		}
	}

	return result
}

// exists tests if a file is accessible.
func exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// mkdir creates a directory if it does not already exist.
func mkdir(name string) error {
	if exists(name) {
		return nil
	}

	return os.Mkdir(name, 0777)
}
