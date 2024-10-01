package main

import (
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
			&cli.StringSliceFlag{
				Name:    "input",
				Aliases: []string{"in", "i"},
				Usage:   "download resources from input `FILE`",
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
		Action: func(c *cli.Context) error {
			references := c.StringSlice("reference")
			inputs := c.StringSlice("input")

			downloader := c.String("downloader")
			output := c.String("output")
			format := c.String("format")

			referenceSet, err := resource.ParseFiles(references)
			if err != nil {
				return err
			}

			inputSet, err := resource.ParseFiles(inputs)
			if err != nil {
				return err
			}

			referenceSet.AddAll(inputSet)

			dl := dler.CreateDownloader(downloader, format)

			dl.SetOutputDirectory(output)

			if _, err := os.Stat(output); err != nil {
				if err := os.Mkdir(output, 0777); err != nil {
					return err
				}
			}

			for _, primaryKey := range inputSet.PrimaryKeys() {
				bestResource := referenceSet.Best(primaryKey)

				err := dl.Download(bestResource)
				if err != nil {
					return err
				}
				log.Println(dl.OutputLocation(bestResource))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
