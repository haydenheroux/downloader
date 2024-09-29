package resource

// primaryKey represents the primary key of a resource.
type primaryKey string

type Resource interface {
	// PrimaryKey returns the unique identifier for the media asset.
	// Resources with the same primary key represent the same media asset.
	PrimaryKey() primaryKey

	// Source returns where the media asset can be accessed.
	Source() string

	// MetadataFields returns the number of metadata fields associated with the media asset.
	MetadataFields() int

	// Title returns the title describing the media asset.
	Title() string
}

type ResourceSet struct {
	// exists tracks whether a resource is contained in the resource set.
	exists map[primaryKey]bool
	// resources maps each primary key to the resources with the same primary key.
	resources map[primaryKey][]Resource
}

// CreateSet creates a new resource set.
func CreateSet(resources []Resource) ResourceSet {
	rs := ResourceSet{
		exists:    make(map[primaryKey]bool),
		resources: make(map[primaryKey][]Resource),
	}

	for _, resource := range resources {
		rs.Add(resource)
	}

	return rs
}

// Add adds a resource to a resource set.
func (rs ResourceSet) Add(resource Resource) {
	key := resource.PrimaryKey()

	rs.exists[key] = true

	_, exists := rs.resources[key]

	if exists == false {
		rs.resources[key] = make([]Resource, 0)
	}

	rs.resources[key] = append(rs.resources[key], resource)
}

// Remove removes a resource from a resource set.
func (rs ResourceSet) Remove(resource Resource) {
	key := resource.PrimaryKey()

	delete(rs.exists, key)
	delete(rs.resources, key)
}

// Contains returns true if the resource set contains the resource.
func (rs ResourceSet) Contains(resource Resource) bool {
	key := resource.PrimaryKey()

	_, exists := rs.exists[key]

	return exists
}

// Resources returns a slice of all resources in the resource set.
func (rs ResourceSet) Resources() []Resource {
	resources := make([]Resource, 0)

	for _, slice := range rs.resources {
		for _, resource := range slice {
			resources = append(resources, resource)
		}
	}

	return resources
}

// Without removes all resources shared with another resource set.
func (rs ResourceSet) Without(other ResourceSet) {
	for _, resource := range rs.Resources() {
		if other.Contains(resource) {
			rs.Remove(resource)
		}
	}
}

// Reduce reduces the resource set such that each primary key is associated with only one resource.
func (rs ResourceSet) Reduce() {
	for key, resources := range rs.resources {
		if len(resources) == 1 {
			continue
		}

		best := pick(resources)

		rs.resources[key] = []Resource{best}
	}
}

// pick returns the resource with the most metadata fields.
func pick(resources []Resource) Resource {
	best := resources[0]

	for _, resource := range resources[1:] {
		if resource.MetadataFields() > best.MetadataFields() {
			best = resource
		}
	}

	return best
}
