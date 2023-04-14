package main

import (
	"errors"
	"net/url"
	"strings"

	"github.com/haydenheroux/cleanstring"
)

type TrackInfo struct {
	URL     string
	Artists string
	Title   string
}

func toURL(s string) string {
	u, err := url.Parse(s)
	isUrl := err == nil && u.Scheme != "" && u.Host != ""

	if isUrl {
		return s
	}

	// Assume the stub is a YouTube ID if it is not a URL
	return "https://www.youtube.com/watch?v=" + s
}

func parse(s string) (TrackInfo, error) {
	fields := strings.Split(s, "\t")
	if len(fields) == 3 {
		track := TrackInfo{
			URL:     toURL(fields[0]),
			Artists: fields[1],
			Title:   fields[2],
		}
		return track, nil
	}
	return TrackInfo{}, errors.New("") // TODO
}

func (t TrackInfo) String() string {
	artists := strings.Split(t.Artists, "&")
	artistString := cleanstring.CleanSlice(artists)

	trackString := cleanstring.Clean(t.Title)

	artistTrack := map[string]string{artistString: trackString}

	return cleanstring.CleanMap(artistTrack)
}
