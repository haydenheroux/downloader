package main

import (
	"music_dl/downloader"
	"music_dl/track"

	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	downloaderExecutable string
	outputExtension      string
	outputFormat         string
	inputFile            string
	outputDirectory      string
)

func init() {
	flag.StringVar(&downloaderExecutable, "d", "yt-dlp", "Name of the executable to be used for downloading tracks. Only downloaders that use youtube-dl-compatible options are supported.")
	flag.StringVar(&outputExtension, "e", "mp3", "Output audio format extension. Used for detecting downloaded files output.")
	flag.StringVar(&outputFormat, "f", "mp3", "Output audio format. Used for specifying to the downloader which format to download.")
	flag.StringVar(&inputFile, "i", "", "Input file. The file must contain lines with three tab-separated (TSV) fields, in this order: URL Artist(s) Title. Multiple artists can be included by delimiting with ampersands (&).")
	flag.StringVar(&outputDirectory, "o", "", "Output directory. If the directory does not exist, it will be created. If this option is not specified, the current working directory is used.")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, "music_dl: ", 0)

	ytdl := downloader.YoutubeDLCompatibleDownloader{
		Executable:      downloaderExecutable,
		Format:          outputFormat,
		FormatExtension: outputFormat,
	}

	file, err := os.Open(inputFile)
	defer file.Close()
	if err != nil {
		logger.Fatalf("Failed to open the input file \"%s\"; error was: %v\n", inputFile, err)
	}

	if err := mkdir(outputDirectory); err != nil {
		logger.Fatalf("Failed to create the output directory \"%s\"; error was: %v\n", outputDirectory, err)
	}

	tracks, err := track.ParseFile(file)
	if err != nil {
		logger.Fatalf("Failed to parse input file; got %d before failing; error was: %v\n", len(tracks), err)
	}

	for _, track := range tracks {
		destination := filepath.Join(outputDirectory, track.String())
		if exists(destination) {
			logger.Printf("Skipping %s because %s already exists\n", track, destination)
			continue
		}
		err := downloader.DownloadTo(ytdl, track, outputDirectory)
		if err != nil {
			logger.Fatalf("Failed to download %s because %v\n", destination, err)
		}
		logger.Printf("Successfully downloaded %s\n", destination)
	}
}
