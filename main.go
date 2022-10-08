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
	INPUT          string
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
	downloadedPath, err := GetDownloadedFilePath(DOWNLOAD_EXECUTABLE, track.URL)
	if err != nil {
		return err
	}
	downloadedPath = ChangeExtension(downloadedPath, OUT_FMT)
	trackName := GetOutputFile(track.Artists, track.Title)
	outPath := filepath.Join(OUT_DIR, trackName)
	err = Move(downloadedPath, outPath)
  return err
}

func printError(scope string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", scope, err)
}

func main() {
	flag.Parse()

	if INPUT == "" {
		printError("args", errors.New("did not specify an input file"))
		os.Exit(1)
	}

	if err := CreateDir(OUT_DIR); err != nil {
		printError("os.Mkdir", err)
		os.Exit(1)
	}

	for lineNum, line := range GetLines(INPUT) {
		track, _ := TrackFrom(line)
    fmt.Println(track)
    trackName := GetOutputFile(track.Artists, track.Title)
    outPath := filepath.Join(OUT_DIR, trackName)
		if Exists(outPath) == false {
      err := download(track)
      if err != nil {
        scope := fmt.Sprintf("%s:%d:%s", INPUT, lineNum, outPath)
        printError(scope, err)
      }
		}
	}
}
