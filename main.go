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

func download(track Track) error {
	err := Download(DOWNLOAD_EXECUTABLE, OUT_FMT, track.URL)
	if err != nil {
		return err
	}
	downloadedFilename, err := DownloadedFilename(DOWNLOAD_EXECUTABLE, track.URL)
	if err != nil {
		return err
	}
	downloadedFilename = ChangeExtension(downloadedFilename, OUT_FMT)
	output := OutputFilename(track.Artists, track.Title)
	outPath := filepath.Join(OUT_DIR, output)
	err = Move(downloadedFilename, outPath)
	return err
}

func printError(scope string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", scope, err)
}

func usage() {
	printError("usage", errors.New("music_dl -i file -o dir"))
}

func main() {
	flag.Parse()

	if INPUT == "" {
		usage()
		os.Exit(1)
	}

	if err := CreateDir(OUT_DIR); err != nil {
		printError("os.Mkdir", err)
		os.Exit(1)
	}

	for lineNum, line := range Lines(INPUT) {
		track, _ := TrackFrom(line)
		output := OutputFilename(track.Artists, track.Title)
		outPath := filepath.Join(OUT_DIR, output)
		if Exists(outPath) == false {
			err := download(track)
			if err != nil {
				scope := fmt.Sprintf("%s:%d:%s", INPUT, lineNum, outPath)
				printError(scope, err)
			}
		}
	}
}
