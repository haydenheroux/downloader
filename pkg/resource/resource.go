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
	// resources maps each primary key to the resources with the same primary key.
	resources map[primaryKey][]Resource
}

// CreateSet creates a new resource set.
func CreateSet(resources []Resource) ResourceSet {
	rs := ResourceSet{
		resources: make(map[primaryKey][]Resource),
	}

	for _, resource := range resources {
		rs.Add(resource)
	}

	return rs
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

// Add adds a resource to a resource set.
func (rs ResourceSet) Add(resource Resource) {
	key := resource.PrimaryKey()

	if _, exists := rs.resources[key]; !exists {
		rs.resources[key] = make([]Resource, 0)
	}

	rs.resources[key] = append(rs.resources[key], resource)
}

// AddAll adds all resources in another set.
func (rs ResourceSet) AddAll(other ResourceSet) {
	for _, resource := range other.Resources() {
		rs.Add(resource)
	}
}

// Remove removes a resource from a resource set.
func (rs ResourceSet) Remove(resource Resource) {
	delete(rs.resources, resource.PrimaryKey())
}

// Contains returns true if the resource set contains the resource.
func (rs ResourceSet) Contains(resource Resource) bool {
	_, exists := rs.resources[resource.PrimaryKey()]

	return exists
}

// PrimaryKeys returns a slice of all primary keys.
func (rs ResourceSet) PrimaryKeys() []primaryKey {
	primaryKeys := make([]primaryKey, 0, len(rs.resources))

	for primaryKey := range rs.resources {
		primaryKeys = append(primaryKeys, primaryKey)
	}

	return primaryKeys
}

// Best returns the best resource (by most metadata) that matches the primary key.
func (rs ResourceSet) Best(primaryKey primaryKey) Resource {
	resources := rs.resources[primaryKey]

	best := resources[0]

	for _, resource := range resources[1:] {
		if resource.MetadataFields() > best.MetadataFields() {
			best = resource
		}
	}

	return best
}
