package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	downloadExecutable string
	goRoutineCount     int
	inputPath          string
	isSilent           bool
	outDir             string
	outExt             bool
	outFmt             string
)

func init() {
	flag.StringVar(&downloadExecutable, "x", "yt-dlp", "youtube-dl executable, yt-dlp supported")
	flag.IntVar(&goRoutineCount, "n", runtime.NumCPU()+2, "maximum number of goroutines")
	flag.StringVar(&inputPath, "i", "", "input file")
	flag.BoolVar(&isSilent, "s", true, "skip printing status messages")
	flag.StringVar(&outDir, "o", "out", "output directory")
	flag.BoolVar(&outExt, "e", true, "append audio format to output filename")
	flag.StringVar(&outFmt, "f", "mp3", "output audio format")
}

func worker(wg *sync.WaitGroup, track TrackInfo) {
	defer wg.Done()
	trackName := CreateOutputFileName(track.Artists, track.Title)
	outPath := filepath.Join(outDir, trackName)
	err := Download(downloadExecutable, outFmt, track.URL)
	if err != nil {
		fmt.Println(track)
		return
	}
	downloadedPath, err := FindOutputFileName(downloadExecutable, track.URL)
	if err != nil {
		return
	}
	downloadedPath = ChangeExtension(downloadedPath, outFmt)
	err = Move(downloadedPath, outPath)
	if err == nil {
		printInfo(fmt.Sprintf("successfully downloaded: %s\n", outPath))
	}
}

func printInfo(msg string) {
	if !isSilent {
		fmt.Printf(msg)
	}
}

func printError(scope string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", scope, err)
}

func main() {
	flag.Parse()

	if inputPath == "" {
		printError("args", errors.New("did not specify an input file"))
		os.Exit(1)
	}

	if err := CreateDir(outDir); err != nil {
		printError("os.Mkdir", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	lines := GetLines(inputPath)
	var tracks []TrackInfo
	for _, line := range lines {
		track, err := GetTrack(line)
		if err != nil {
			printError("GetTrack", err)
		}
		trackName := CreateOutputFileName(track.Artists, track.Title)
		outPath := filepath.Join(outDir, trackName)
		if Exists(outPath) {
			printInfo(fmt.Sprintf("%s: already exists, not downloading\n", outPath))
		} else {
			tracks = append(tracks, track)
		}
	}
	chunks := Chunkify(tracks, goRoutineCount)
	for _, chunk := range chunks {
		for _, track := range chunk {
			wg.Add(1)
			go worker(&wg, track)
		}
		wg.Wait()
	}

}
