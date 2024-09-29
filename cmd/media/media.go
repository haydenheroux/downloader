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

	if shouldMkdir() {
		if err := mkdir(outputDirectory); err != nil {
			logger.Fatalf("failed to make output directory (%v)\n", err)
		}
	}

	set, err := parseFiles(files)

	if err != nil {
		logger.Fatalf("failed to parse input file (%v)", err)
	}

	if unique {
		set.Unique()
	}

	if listResources {
		for _, resource := range set.Resources() {
			fmt.Println(resource.Name())
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

func shouldMkdir() bool {
	return outputDirectory != defaultOutputDir
}

func parseFiles(names []string) (resource.ResourceSet, error) {
	result := resource.CreateSet([]resource.Resource{})

	for _, name := range names {
		resources, err := parseFile(name)

		if err != nil {
			return resource.ResourceSet{}, err
		}

		for _, resource := range resources.Resources() {
			result.Add(resource)
		}
	}

	return result, nil
}

func parseFile(name string) (resource.ResourceSet, error) {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		return resource.ResourceSet{}, err
	}

	resources, err := resource.ParseFile(file)
	if err != nil {
		return resource.ResourceSet{}, err
	}

	return resources, nil
}

// exists tests if a file is accessible.
func exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// mkdir creates a directory if it does not already exist.
func mkdir(name string) error {
	if exists(name) {
		return nil
	}

	return os.Mkdir(name, 0777)
}
