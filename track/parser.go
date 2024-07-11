package track

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/haydenheroux/strfmt"
)

func parseURLName(url, name string) Track {
	return Track{
		URL:  url,
		Name: name,
	}
}

func parseURLInfo(url, artists, title string) Track {
	A := strfmt.Join(strings.Split(artists, "&"))
	name := strfmt.Associate(map[string]string{A: title})

	return parseURLName(url, name)
}

func parse(fields []string) Track {
	switch len(fields) {
	case 2:
		return parseURLName(fields[0], fields[1])
	case 3:
		return parseURLInfo(fields[0], fields[1], fields[2])
	default:
		return Track{}
	}
}

func ParseFile(r io.Reader) ([]Track, error) {
	reader := csv.NewReader(r)
	// Allow variable number of fields
	reader.FieldsPerRecord = -1

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
