package track

import (
	"encoding/csv"
	"io"
	"net/url"
	"strings"
)

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

func parse(fields []string) Track {
	track := Track{
		URL:     toURL(fields[0]),
		Artists: strings.Split(fields[1], "&"),
		Title:   fields[2],
	}
	return track
}

func ParseFile(r io.Reader) ([]Track, error) {
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	reader.FieldsPerRecord = 3

	records, err := reader.ReadAll()
	if err != nil {
		return []Track{}, err
	}

	tracks := make([]Track, 0, len(records))

	for _, record := range records {
		tracks = append(tracks, parse(record))
	}

	return tracks, nil
}
