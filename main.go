package main

import (
	"fmt"
	"music_dl/downloader"
	"music_dl/track"

	"flag"
	"log"
	"os"
)

var (
	downloaderName  string
	outputFormat    string
	inputFile       string
	outputDirectory string
	printInfo       bool
)

func init() {
	flag.StringVar(&downloaderName, "d", "ytdl", "Name of the downloader to be used for downloading tracks.")
	flag.StringVar(&outputFormat, "f", "mp3", "Output audio format. Used for specifying to the downloader which format to download.")
	flag.StringVar(&inputFile, "i", "", "Input file. The file must contain lines with three tab-separated (TSV) fields, in this order: URL Artist(s) Title. Multiple artists can be included by delimiting with ampersands (&).")
	flag.StringVar(&outputDirectory, "o", "", "Output directory. If the directory does not exist, it will be created. If this option is not specified, the current working directory is used.")
	flag.BoolVar(&printInfo, "p", false, "Print info. If true, print additional information about each downloaded track.")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, "music_dl: ", 0)

	dl := downloader.CreateDownloader(downloaderName, outputFormat)
	if dl == nil {
		logger.Fatalf("Failed to initialize downloader; name was %s\n", downloaderName)
	}

	file, err := os.Open(inputFile)
	defer file.Close()

	if err != nil {
		logger.Fatalf("Failed to open the input file \"%s\": %v\n", inputFile, err)
	}

    if outputDirectory != "" {
        if err := mkdir(outputDirectory); err != nil {
            logger.Fatalf("Failed to create the output directory \"%s\": %v\n", outputDirectory, err)
        }
    }

	tracks, err := track.ParseFile(file)
	if err != nil {
		logger.Fatalf("Failed to parse input file; got %d before failing: %v\n", len(tracks), err)
	}

	tracks = removeExisting(tracks, dl, outputDirectory)

	for _, track := range tracks {
		if printInfo {
			fmt.Printf("%s\n", track)
		}

		err := dl.Download(track, outputDirectory)
		if err != nil {
			logger.Fatalf("Failed to download %s: %v\n", track, err)
		}
	}
}
