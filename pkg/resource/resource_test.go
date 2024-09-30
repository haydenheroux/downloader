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
		t.Error("Set does not match slice")
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
		t.Error("Set before adding matches slice after adding")
	}

	set.Add(newResource)

	if !exactMatch(set.Resources(), testResources) {
		t.Error("Set after adding does not match slice after adding")
	}
}

func removeResource(resources []Resource, remove Resource) []Resource {
	var removedResources []Resource

	for _, resource := range resources {
		if resource.PrimaryKey() != remove.PrimaryKey() {
			removedResources = append(removedResources, resource)
		}
	}

	return removedResources
}

func TestRemove(t *testing.T) {
	testResources := testResources()

	resource := attributedUrl{
		url:      "",
		creators: []string{"Frank Herbert"},
		name:     "Dune",
	}

	set := CreateSet(testResources)

	removedResources := removeResource(testResources, resource)

	if exactMatch(set.Resources(), removedResources) {
		t.Error("Set before removal matches slice after removal")
	}

	set.Remove(resource)

	if exactMatch(set.Resources(), testResources) {
		t.Error("Set after removal matches slice before removal")
	}

	if !exactMatch(set.Resources(), removedResources) {
		t.Error("Set after removal does not match slice after removal")
	}
}
