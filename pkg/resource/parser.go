package resource

import (
	"encoding/csv"
	"io"
)

func parse(fields []string) Resource {
	switch len(fields) {
	case 1:
		return namedUrl{
			url:  fields[0],
			name: fields[0],
		}
	case 2:
		return namedUrl{
			url:  fields[0],
			name: fields[1],
		}
	case 3:
		return createAttributedURL(fields)
	default:
		resource := createAttributedURL(fields)

		tags := fields[3:]

		return taggedResource{
			resource: resource,
			tags:     tags,
		}
	}
}

func ParseFile(r io.Reader) (ResourceSet, error) {
	reader := csv.NewReader(r)
	// Allow variable number of fields
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return ResourceSet{}, err
	}

	resources := make([]Resource, 0, len(records))

	for _, record := range records {
		resources = append(resources, parse(record))
	}

	return CreateSet(resources), nil
}
