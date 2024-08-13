package main

import (
	"downloader/downloader"
	"downloader/resource"
	"fmt"

	"flag"
	"log"
	"os"
)

const (
	APP_NAME                 = "downloader"
	DEFAULT_OUTPUT_DIRECTORY = ""
)

var (
	downloaderName  string
	outputFormat    string
	outputDirectory string

	listTracks bool
	printInfo  bool
)

func init() {
	flag.StringVar(&downloaderName, "d", "ytdl", "downloader name")
	flag.StringVar(&outputFormat, "f", "mp3", "output format")
	flag.StringVar(&outputDirectory, "o", DEFAULT_OUTPUT_DIRECTORY, "output directory")

	flag.BoolVar(&printInfo, "p", false, "print information as a track is downloading")
	flag.BoolVar(&listTracks, "l", false, "list tracks that would be downloaded then exit")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, APP_NAME+": ", 0)

	files := flag.Args()

	if len(files) == 0 {
		logger.Fatalf("no input files provided\n")
	}

	dl := downloader.CreateDownloader(downloaderName, outputFormat)

	if shouldMkdir() {
		if err := mkdir(outputDirectory); err != nil {
			logger.Fatalf("failed to make output directory (%v)\n", err)
		}
	}

	tracks, err := parseFiles(files)

	if err != nil {
		logger.Fatalf("failed to parse input file (%v)", err)
	}

	if listTracks {
		for _, track := range tracks {
			fmt.Println(track)
		}

		os.Exit(0)
	}

	existing := existingResources(tracks, dl, outputDirectory)

	if printInfo {
		for track := range existing {
			fmt.Printf("found: %s\n", track)
		}
	}

	tracks = resource.Difference(tracks, existing)
	tracks = resource.Unique(tracks)

	for _, track := range tracks {
		if printInfo {
			fmt.Printf("started: %s\n", track)
		}

		err := dl.Download(track, outputDirectory)
		if err != nil {
			logger.Fatalf("failed to download %s (%v)\n", track, err)
		}

		if printInfo {
			fmt.Printf("completed: %s\n", track)
		}
	}
}

func shouldMkdir() bool {
	return outputDirectory != DEFAULT_OUTPUT_DIRECTORY
}

func parseFiles(names []string) ([]resource.Resource, error) {
	result := make([]resource.Resource, 0)

	for _, name := range names {
		tracks, err := parseFile(name)

		if err != nil {
			return result, err
		}

		for _, track := range tracks {
			result = append(result, track)
		}
	}

	return result, nil
}

func parseFile(name string) ([]resource.Resource, error) {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		return []resource.Resource{}, err
	}

	tracks, err := resource.ParseFile(file)
	if err != nil {
		return []resource.Resource{}, err
	}

	return tracks, nil
}
