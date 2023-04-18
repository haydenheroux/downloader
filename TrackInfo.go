package main

import (
	"errors"
	"net/url"
	"strings"

	"github.com/haydenheroux/cleanstring"
)

// TrackInfo represents a track.
type TrackInfo struct {
	// URL is the URL required to locate the track for download.
	URL     string
	// Artists is a list of artists or creators of the track.
	Artists []string
	// Title is the title of the track.
	Title   string
}

// toURL transforms a string to a valid URL.
// If s is already a URL, this function is transparent.
// Otherwise, s is assumed to be the ID field of a YouTube URL.
func toURL(s string) string {
	u, err := url.Parse(s)
	isUrl := err == nil && u.Scheme != "" && u.Host != ""

	if isUrl {
		return s
	}

	// Assume the stub is a YouTube ID if it is not a URL
	return "https://www.youtube.com/watch?v=" + s
}

// parse attempts to create a TrackInfo struct from an input string.
func parse(s string) (TrackInfo, error) {
	fields := strings.Split(s, "\t")

	if len(fields) != 3 {
		return TrackInfo{}, errors.New("") // TODO
	}

	track := TrackInfo{
		URL:     toURL(fields[0]),
		Artists: strings.Split(fields[1], "&"),
		Title:   fields[2],
	}
	return track, nil
}

// String returns the representation of this track as a string.
func (t TrackInfo) String() string {
	artists := cleanstring.CleanSlice(t.Artists)
	title := cleanstring.Clean(t.Title)
	temp := map[string]string{artists: title}
	return cleanstring.CleanMap(temp)
}
