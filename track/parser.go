package track

import (
	"bufio"
	"errors"
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

func ParseSlice(strs []string) ([]Track, error) {
	tracks := make([]Track, len(strs))
	for n, s := range strs {
		track, err := Parse(s)

		if err != nil {
			return tracks, err
		}

		tracks[n] = track
	}

	return tracks, nil
}

func ParseFile(r io.Reader) ([]Track, error) {
	lines := []string{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []Track{}, err
	}

	return ParseSlice(lines)
}
