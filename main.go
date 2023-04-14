package main

import (
	"errors"
	"flag"
	"fmt"
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

func printError(scope string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", scope, err)
}

func usage() {
	printError("usage", errors.New("music_dl -i file -o dir"))
}


func downloadTo(d Downloader, t TrackInfo, directory string) error {
	err := d.Download(t)
	if err != nil {
		return err
	}
	filename, err := d.GetFilename(t)
	if err != nil {
		return err
	}
	destination := filepath.Join(directory, filename)
	return os.Rename(filename, destination)
}

func main() {
	flag.Parse()

	if INPUT == "" {
		usage()
		os.Exit(1)
	}

	downloader := YoutubeDLCompatibleDownloader{
		Executable: DOWNLOAD_EXECUTABLE,
		Format: OUT_FMT,
		FormatExtension: OUT_FMT,
	}

	if err := CreateDir(OUT_DIR); err != nil {
		printError("os.Mkdir", err)
		os.Exit(1)
	}

	lines := Lines(INPUT)

	tracks := []TrackInfo{}
	
	for n, line := range lines {
		track, _ := Parse(line)
		fmt.Fprintf(os.Stderr, "Successfully parsed line %d: %s\n", n, track)
		tracks = append(tracks, track)
	}

	for _, track := range tracks {
		destination := filepath.Join(OUT_DIR, track.String())
		if Exists(destination) {
			fmt.Fprintf(os.Stderr, "%s exists; skipping %s\n", destination, track)
			continue
		}
		fmt.Println(track)
		err := downloadTo(downloader, track, OUT_DIR)
		if err != nil {
			printError(destination, err)
		}
	}
}
