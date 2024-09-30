package resource

import (
	"encoding/csv"
	"io"
	"os"
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
		return createAttributedUrl(fields)
	default:
		resource := createAttributedUrl(fields)

		tags := fields[3:]

		return taggedResource{
			resource: resource,
			tags:     tags,
		}
	}
}

func parseInput(input io.Reader) (ResourceSet, error) {
	reader := csv.NewReader(input)
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

func ParseFile(name string) (ResourceSet, error) {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		return ResourceSet{}, err
	}

	return parseInput(file)
}

func ParseFiles(names []string) (ResourceSet, error) {
	result := CreateSet([]Resource{})

	for _, name := range names {
		resources, err := ParseFile(name)

		if err != nil {
			return ResourceSet{}, err
		}

		for _, resource := range resources.Resources() {
			result.Add(resource)
		}
	}

	return result, nil
}
