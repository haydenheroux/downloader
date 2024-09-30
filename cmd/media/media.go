package main

import (
	"log"
	"os"

	"github.com/haydenheroux/media/pkg/downloader"
	"github.com/haydenheroux/media/pkg/resource"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "media",
		Usage: "manage and download media assets",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "reference",
				Aliases: []string{"r"},
				Usage:   "Use `FILE` as reference",
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "download",
				Usage:     "download media assets",
				ArgsUsage: "[files]",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "input",
						Aliases: []string{"i", "in"},
						Usage:   "Download assets in `INPUT`",
					},
					&cli.StringFlag{
						Name:     "downloader",
						Aliases:  []string{"d"},
						Usage:    "Downloader name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "format",
						Aliases:  []string{"f"},
						Usage:    "Output format",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o", "out"},
						Usage:    "Output directory",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					references := c.StringSlice("reference")
					inputs := c.StringSlice("input")

					downloaderName := c.String("downloader")
					outputDirectory := c.String("output")
					outputFormat := c.String("format")

					referenceSet, err := resource.ParseFiles(references)
					if err != nil {
						return err
					}

					inputSet, err := resource.ParseFiles(inputs)
					if err != nil {
						return err
					}

					referenceSet.AddAll(inputSet)

					dl := downloader.CreateDownloader(downloaderName, outputFormat)

					if _, err := os.Stat(outputDirectory); err != nil {
						if err := os.Mkdir(outputDirectory, 0777); err != nil {
							return err
						}
					}

					for _, primaryKey := range inputSet.PrimaryKeys() {
						bestResource := referenceSet.Best(primaryKey)

						log.Printf("starting %s\n", bestResource.Title())
						err := dl.Download(bestResource, outputDirectory)
						if err != nil {
							return err
						}
						log.Printf("finished %s\n", bestResource.Title())
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
