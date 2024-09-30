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

func containsAll(set ResourceSet, resources []Resource) bool {
	resourcesCopy := make([]Resource, len(resources))
	copy(resourcesCopy, resources)

	if len(set.Resources()) != len(resources) {
		return false
	}

	for _, setResource := range set.Resources() {
		found := false

		for index, copyResource := range resourcesCopy {
			if setResource.PrimaryKey() == copyResource.PrimaryKey() {
				// Instead of deleting the resource, replace with a canary resource
				resourcesCopy[index] = canaryResource{}
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	for _, resource := range resourcesCopy {
		if resource.PrimaryKey() != canaryValue {
			return false
		}
	}

	return true
}

func TestCreateSet(t *testing.T) {
	testResources := testResources()

	set := CreateSet(testResources)

	if !containsAll(set, testResources) {
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

	set.Add(newResource)

	if !containsAll(set, testResources) {
		t.Error("Not all resources from the source slice were added to the resource set")
	}
}
