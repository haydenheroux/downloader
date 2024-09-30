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

func phonyResource() Resource {
	return namedUrl{
		url:  "",
		name: "",
	}
}

func TestCreateSet(t *testing.T) {
	set := CreateSet(testResources())

	testResources := testResources()

	for _, resource := range set.Resources() {
		found := false

		for index, testResource := range testResources {
			if resource != nil && resource.PrimaryKey() == testResource.PrimaryKey() {
				// Instead of deleting the resource, replace with a phony resource
				testResources[index] = phonyResource()
				found = true
				break
			}
		}

		if !found {
			t.Error("Not all resources from the source slice were added to the resource set")
		}
	}
}
