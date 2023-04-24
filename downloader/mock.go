package downloader

import (
	"music_dl/track"
)

type MockDownloader struct {
	Format string
	OutputDirectory string
}

func (m MockDownloader) Download(track track.Track) error {
	return nil
}

func (m MockDownloader) GetOutputFilename(track track.Track) string {
	return track.String() + "." + m.Format
}
