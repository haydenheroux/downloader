package main

import (
	"fmt"
	"downloader/downloader"
	"downloader/track"

	"flag"
	"log"
	"os"
)

const (
	APP_NAME                 = "downloader"
	NO_INPUT_FILE            = ""
	DEFAULT_OUTPUT_DIRECTORY = ""
)

var (
	downloaderName  string
	outputFormat    string
	inputFile       string
	outputDirectory string
	printInfo       bool
)

func init() {
	flag.StringVar(&downloaderName, "d", "ytdl", "Name of the downloader used to download tracks.")
	flag.StringVar(&outputFormat, "f", "mp3", "Output format. Usually the extension of the output file.")
	flag.StringVar(&inputFile, "i", NO_INPUT_FILE, "Input file. The file must contain lines with three tab-separated (TSV) fields, in this order: URL Artist(s) Title. Multiple artists can be included by delimiting with ampersands (&).")
	flag.StringVar(&outputDirectory, "o", DEFAULT_OUTPUT_DIRECTORY, "Output directory. If the directory does not exist, it will be created.")
	flag.BoolVar(&printInfo, "p", false, "Print info. If true, print additional information about each downloaded track.")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, APP_NAME+": ", 0)

	if inputFile == NO_INPUT_FILE {
		logger.Fatalf("No input file provided\n")
	}

	dl := downloader.CreateDownloader(downloaderName, outputFormat)

	if shouldMkdir() {
		if err := mkdir(outputDirectory); err != nil {
			logger.Fatalf("Failed to make output directory (%v)\n", err)
		}
	}

	tracks, err := parseTracks()
	if err != nil {
		logger.Fatalf("Failed to parse tracks (%v)\n", err)
	}

	tracks = removeExisting(tracks, dl, outputDirectory)

	for _, track := range tracks {
		if printInfo {
			fmt.Printf("%s\n", track)
		}

		err := dl.Download(track, outputDirectory)
		if err != nil {
			logger.Fatalf("Failed to download %s (%v)\n", track, err)
		}
	}
}

func shouldMkdir() bool {
	return outputDirectory != DEFAULT_OUTPUT_DIRECTORY
}

func parseTracks() ([]track.Track, error) {
	file, err := os.Open(inputFile)
	defer file.Close()

	if err != nil {
		return []track.Track{}, err
	}

	tracks, err := track.ParseFile(file)
	if err != nil {
		return []track.Track{}, err
	}

	return tracks, nil
}
