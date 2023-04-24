package downloader

import (
	"music_dl/track"

	"os"
)

type MockDownloader struct {
	Format string
	OutputDirectory string
}

func (m MockDownloader) Download(track track.Track) error {
	name := m.GetOutputFilename(track)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		defer file.Close()

		return err
	}
	return nil
}

func (m MockDownloader) GetOutputFilename(track track.Track) string {
	return track.String() + "." + m.Format
}
