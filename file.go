package main

import "music_dl/track"
import "path/filepath"
import "os"

// onlyMissingFrom returns the tracks that are missing from the directory.
func onlyMissingFrom(tracks []track.Track, directory string) []track.Track {
	result := make([]track.Track, 0, len(tracks))

	for _, track := range tracks {
		if out := filepath.Join(directory, track.String()); !exists(out) {
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
