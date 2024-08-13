package resource

import (
	"encoding/csv"
	"io"
	"strings"
)

func parse(fields []string) Resource {
	switch len(fields) {
	case 2:
		return urlName{
			url:  fields[0],
			name: fields[1],
		}
	case 3:
		return urlArtistsTitle{
			url:     fields[0],
			artists: strings.Split(fields[1], "&"),
			title:   fields[2],
		}
	default:
		return urlName{}
	}
}

func ParseFile(r io.Reader) ([]Resource, error) {
	reader := csv.NewReader(r)
	// Allow variable number of fields
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return []Resource{}, err
	}

	tracks := make([]Resource, 0, len(records))

	for _, record := range records {
		tracks = append(tracks, parse(record))
	}

	return tracks, nil
}
