package track

import (
	"errors"
	"net/url"
	"strings"

	"github.com/haydenheroux/cleanstring"
)

// Track represents a track.
type Track struct {
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

// Parse attempts to create a Track struct from an input string.
func Parse(s string) (Track, error) {
	fields := strings.Split(s, "\t")

	if len(fields) != 3 {
		return Track{}, errors.New("") // TODO
	}

	track := Track{
		URL:     toURL(fields[0]),
		Artists: strings.Split(fields[1], "&"),
		Title:   fields[2],
	}
	return track, nil
}

// String returns the representation of this track as a string.
func (t Track) String() string {
	artists := cleanstring.CleanSlice(t.Artists)
	title := cleanstring.Clean(t.Title)
	temp := map[string]string{artists: title}
	return cleanstring.CleanMap(temp)
}
