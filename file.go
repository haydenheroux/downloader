package main

import (
	"downloader/downloader"
	"downloader/resource"
	"os"
)

// scanExisting computes the set of tracks existing in a directory.
func scanExisting(tracks []resource.Resource, downloader downloader.Downloader, directory string) map[resource.Resource]bool {
	existing := make(map[resource.Resource]bool)

	for _, track := range tracks {
		if exists(downloader.GetOutputFilename(track, directory)) {
			existing[track] = true
		}
	}

	return existing
}

// removeExisting removes existing tracks from a slice of tracks.
func removeExisting(tracks []resource.Resource, existing map[resource.Resource]bool) []resource.Resource {
	result := make([]resource.Resource, 0, len(tracks))

	for _, track := range tracks {
		if exists, _ := existing[track]; !exists {
			result = append(result, track)
		}
	}

	return result
}

// removeDuplicates removes tracks that have the same name.
func removeDuplicates(tracks []resource.Resource) []resource.Resource {
	result := make([]resource.Resource, 0, len(tracks))

	set := make(map[string]bool)

	for _, track := range tracks {
		name := track.Name()

		if exists, _ := set[name]; !exists {
			result = append(result, track)
			set[name] = true
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
