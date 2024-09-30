package resource

import (
	"testing"
)

func testResources() []Resource {
	return []Resource{
		attributedUrl{
			url:      "",
			creators: []string{"George Orwell"},
			name:     "1984",
		},
		attributedUrl{
			url:      "",
			creators: []string{"Frank Herbert"},
			name:     "Dune",
		},
		taggedResource{
			resource: attributedUrl{
				url:      "",
				creators: []string{"Frank Herbert"},
				name:     "Dune",
			},
			tags: []string{"science fiction", "book"},
		},
		attributedUrl{
			url:      "",
			creators: []string{"Orson Scott Card"},
			name:     "Ender's Game",
		},
	}
}

type canaryResource struct{}

const canaryValue = "canary"

func (c canaryResource) PrimaryKey() primaryKey {
	return canaryValue
}

func (c canaryResource) Source() string {
	return canaryValue
}

func (c canaryResource) MetadataFields() int {
	return 0
}

func (c canaryResource) Title() string {
	return canaryValue
}

func exactMatch(a, b []Resource) bool {
	if len(a) != len(b) {
		return false
	}

	countMap := make(map[primaryKey]int)

	for _, resource := range a {
		countMap[resource.PrimaryKey()]++
	}

	for _, resource := range b {
		countMap[resource.PrimaryKey()]--
		if countMap[resource.PrimaryKey()] < 0 {
			return false
		}
	}

	for _, count := range countMap {
		if count != 0 {
			return false
		}
	}

	return true
}

func TestCreateSet(t *testing.T) {
	testResources := testResources()

	set := CreateSet(testResources)

	if !exactMatch(set.Resources(), testResources) {
		t.Error("Not all resources from the source slice were added to the resource set")
	}
}

func TestAdd(t *testing.T) {
	testResources := testResources()

	newResource := attributedUrl{
		url:      "",
		creators: []string{"Ray Bradbury"},
		name:     "Fahrenheit 451",
	}

	set := CreateSet(testResources)

	testResources = append(testResources, newResource)

	if exactMatch(set.Resources(), testResources) {
		t.Error("Set contains all resources before they were added")
	}

	set.Add(newResource)

	if !exactMatch(set.Resources(), testResources) {
		t.Error("Not all resources from the source slice were added to the resource set")
	}
}
