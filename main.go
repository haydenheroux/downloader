package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	downloaderExecutable string
	outputExtension             string
	outputFormat             string
	inputFile               string
	outputDirectory             string
)

func init() {
	flag.StringVar(&downloaderExecutable, "d", "yt-dlp", "Name of the executable to be used for downloading tracks. Only downloaders that use youtube-dl-compatible options are supported.")
	flag.StringVar(&outputExtension, "e", "mp3", "Output audio format extension. Used for detecting downloaded files output.")
	flag.StringVar(&outputFormat, "f", "mp3", "Output audio format. Used for specifying to the downloader which format to download.")
	flag.StringVar(&inputFile, "i", "", "Input file. The file must contain lines with three tab-separated (TSV) fields, in this order: URL Artist(s) Title. Multiple artists can be included by delimiting with ampersands (&).")
	flag.StringVar(&outputDirectory, "o", "", "Output directory. If the directory does not exist, it will be created. If this option is not specified, the current working directory is used.")
}

func downloadTo(downloader Downloader, track TrackInfo, directory string) error {
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

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, "music_dl: ", 0)

	if inputFile == "" {
		logger.Fatalln("Usage: music_dl -i inputFile")
	}

	downloader := YoutubeDLCompatibleDownloader{
		Executable:      downloaderExecutable,
		Format:          outputFormat,
		FormatExtension: outputFormat,
	}

	if err := mkdir(outputDirectory); err != nil {
		logger.Fatalln("Creating directory")
	}

	tracks := []TrackInfo{}

	lines, _ := linesIn(inputFile)

	for n, line := range lines {
		track, _ := parse(line)
		logger.Printf("Successfully parsed line %d: %s\n", n, track)
		tracks = append(tracks, track)
	}

	for _, track := range tracks {
		destination := filepath.Join(outputDirectory, track.String())
		if exists(destination) {
			logger.Printf("Skipping %s because %s already exists\n", track, destination)
			continue
		}
		err := downloadTo(downloader, track, outputDirectory)
		if err != nil {
			logger.Fatalf("Failed to download %s because %v\n", destination, err)
		}
		logger.Printf("Successfully downloaded %s\n", destination)
	}
}
