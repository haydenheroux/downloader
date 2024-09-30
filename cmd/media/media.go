package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/haydenheroux/media/pkg/downloader"
	"github.com/haydenheroux/media/pkg/resource"
)

const (
	appName          = "media"
	defaultOutputDir = ""
)

var (
	downloaderName  string
	outputFormat    string
	outputDirectory string

	listResources bool
	printInfo     bool
	unique        bool
)

func init() {
	flag.StringVar(&downloaderName, "d", "ytdl", "downloader name")
	flag.StringVar(&outputFormat, "f", "mp3", "output format")
	flag.StringVar(&outputDirectory, "o", defaultOutputDir, "output directory")

	flag.BoolVar(&printInfo, "p", false, "print information as a resource is downloading")
	flag.BoolVar(&listResources, "l", false, "list resources that would be downloaded then exit")
	flag.BoolVar(&unique, "u", false, "attempt to identify unique resources and remove duplicates before downloading")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, appName+": ", 0)

	files := flag.Args()

	if len(files) == 0 {
		logger.Fatalf("no input files provided\n")
	}

	dl := downloader.CreateDownloader(downloaderName, outputFormat)

	if outputDirectory != defaultOutputDir {
		if _, err := os.Stat(outputDirectory); err != nil {
			if err := os.Mkdir(outputDirectory, 0777); err != nil {
				logger.Fatalf("failed to make output directory (%v)\n", err)
			}
		}
	}

	set, err := resource.ParseFiles(files)

	if err != nil {
		logger.Fatalf("failed to parse input file (%v)", err)
	}

	if unique {
		set.Reduce()
	}

	if listResources {
		for _, resource := range set.Resources() {
			fmt.Println(resource.Title())
		}

		os.Exit(0)
	}

	for _, resource := range set.Resources() {
		name := dl.GetOutputFilename(resource, outputDirectory)

		if printInfo {
			fmt.Printf("started: %s\n", name)
		}

		err := dl.Download(resource, outputDirectory)
		if err != nil {
			logger.Fatalf("failed to download %s (%v)\n", name, err)
		}

		if printInfo {
			fmt.Printf("completed: %s\n", name)
		}
	}
}
