package main

import (
	"fmt"
	"log"
	"os"

	dler "github.com/haydenheroux/media/pkg/downloader"
	"github.com/haydenheroux/media/pkg/resource"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "media",
		Usage: "download media assets",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "reference",
				Aliases: []string{"ref", "r"},
				Usage:   "reference `FILE` for metadata",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "list reference keys",
				Action: listReferenceKeys,
			},
			{
				Name:  "download",
				Usage: "download resources",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "input",
						Aliases: []string{"in", "i"},
						Usage:   "download resources from input `FILE`",
					},
					&cli.BoolFlag{
						Name:    "keys",
						Aliases: []string{"k"},
						Usage:   "input file contains keys",
					},
					&cli.StringFlag{
						Name:     "downloader",
						Aliases:  []string{"dl", "d"},
						Usage:    "download using the downloader named `NAME`",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "format",
						Aliases:  []string{"fmt", "f"},
						Usage:    "output files in the format `FORMAT`",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"out", "o"},
						Usage:    "output files in the directory `DIR`",
						Required: true,
					},
				},
				Action: downloadResources,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func listReferenceKeys(c *cli.Context) error {
	references := c.StringSlice("reference")

	referenceSet, err := resource.ParseFiles(references)
	if err != nil {
		return err
	}

	for _, key := range referenceSet.PrimaryKeys() {
		fmt.Println(key)
	}
	return nil
}

func downloadResources(c *cli.Context) error {
	keys := c.Bool("keys")

	references := c.StringSlice("reference")
	inputs := c.StringSlice("input")

	downloader := c.String("downloader")
	output := c.String("output")
	format := c.String("format")

	referenceSet, err := resource.ParseFiles(references)
	if err != nil {
		return err
	}

	var inputKeys []resource.PrimaryKey

	if keys {
		inputKeys, err = resource.ParseKeyFiles(inputs)
		if err != nil {
			return err
		}
	} else {
		inputSet, err := resource.ParseFiles(inputs)
		if err != nil {
			return err
		}

		referenceSet.AddAll(inputSet)

		inputKeys = inputSet.PrimaryKeys()
	}

	dl := dler.CreateDownloader(downloader, format)

	dl.SetOutputDirectory(output)

	if _, err := os.Stat(output); err != nil {
		if err := os.Mkdir(output, 0777); err != nil {
			return err
		}
	}

	for _, key := range inputKeys {
		if !referenceSet.ContainsKey(key) {
			return fmt.Errorf("key %s does not match any reference", key)
		}

		bestResource := referenceSet.Best(key)

		log.Printf("started %s\n", dl.OutputLocation(bestResource))
		err := dl.Download(bestResource)
		if err != nil {
			return err
		}
		log.Printf("finished %s\n", dl.OutputLocation(bestResource))
	}

	return nil
}
