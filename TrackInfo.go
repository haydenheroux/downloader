package main

import (
	"errors"
	"net/url"
	"strings"
)

type TrackInfo struct {
	URL     string
	Artists string
	Title   string
}

func Chunkify(lines []TrackInfo, size int) [][]TrackInfo {
	var chunks [][]TrackInfo

	for start := 0; start < len(lines); start += size {
		end := start + size
		if start > len(lines) {
			start = len(lines)
		}
		if end > len(lines) {
			end = len(lines)
		}

		chunk := lines[start:end]
		chunks = append(chunks, chunk)
	}

	return chunks
}

func toDownloadURL(urlStub string) string {
	u, err := url.Parse(urlStub)
	isURL := err == nil && u.Scheme != "" && u.Host != ""
	if isURL {
		// Since it already is a URL, just return it
		return urlStub
	}
	return "https://www.youtube.com/watch?v=" + urlStub
}

func GetTrack(line string) (TrackInfo, error) {
	track := TrackInfo{
		URL:     "",
		Artists: "",
		Title:   "",
	}
	fields := strings.Split(line, "\t")
	if len(fields) != 3 {
		return track, errors.New("not enough fields in line")
	}
	track.URL = toDownloadURL(fields[0])
	track.Artists = fields[1]
	track.Title = fields[2]
	return track, nil
}
