package main

import (
	"downloader/downloader"
	"downloader/resource"
	"os"
)

func existingResources(resources []resource.Resource, downloader downloader.Downloader, directory string) map[resource.Resource]bool {
	existing := make(map[resource.Resource]bool)

	for _, track := range resources {
		if exists(downloader.GetOutputFilename(track, directory)) {
			existing[track] = true
		}
	}

	return existing
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
