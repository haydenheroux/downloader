package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	DOWNLOAD_EXECUTABLE string
	INPUT               string
	OUT_DIR             string
	OUT_FMT             string
)

func init() {
	flag.StringVar(&DOWNLOAD_EXECUTABLE, "x", "yt-dlp", "youtube-dl executable, yt-dlp supported")
	flag.StringVar(&INPUT, "i", "", "input file")
	flag.StringVar(&OUT_DIR, "o", "out", "output directory")
	flag.StringVar(&OUT_FMT, "f", "mp3", "output audio format")
}

func downloadTo(downloader Downloader, track TrackInfo, directory string) (string, error) {
	err := downloader.Download(track)
	if err != nil {
		return "", err
	}

	filename, err := downloader.GetFilename(track)
	if err != nil {
		return "", err
	}

	destination := filepath.Join(directory, track.String())
	err = os.Rename(filename, destination)
	if err != nil {
		return "", err
	}

	return destination, nil
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, "music_dl: ", 0)

	if INPUT == "" {
		logger.Fatalln("Usage: music_dl -i INPUT_FILE -o OUTPUT_DIRECTORY")
	}

	downloader := YoutubeDLCompatibleDownloader{
		Executable:      DOWNLOAD_EXECUTABLE,
		Format:          OUT_FMT,
		FormatExtension: OUT_FMT,
	}

	if err := mkdir(OUT_DIR); err != nil {
		logger.Fatalln("Creating directory")
	}

	tracks := []TrackInfo{}

	lines, _ := linesIn(INPUT)

	for n, line := range lines {
		track, _ := Parse(line)
		logger.Printf("Successfully parsed line %d: %s\n", n, track)
		tracks = append(tracks, track)
	}

	for _, track := range tracks {
		destination := filepath.Join(OUT_DIR, track.String())
		if exists(destination) {
			logger.Printf("Skipping %s because %s already exists\n", track, destination)
			continue
		}
		destination, err := downloadTo(downloader, track, OUT_DIR)
		if err != nil {
			logger.Fatalf("Failed to download %s because %v\n", destination, err)
		}
		logger.Printf("Successfully downloaded %s\n", destination)
	}
}
