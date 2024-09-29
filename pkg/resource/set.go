package resource

type ResourceSet struct {
	exists    map[key]bool
	resources map[key][]Resource
}

func CreateSet(resources []Resource) ResourceSet {
	rs := ResourceSet{
		exists:    make(map[key]bool),
		resources: make(map[key][]Resource),
	}

	for _, resource := range resources {
		rs.Add(resource)
	}

	return rs
}

func (rs ResourceSet) Add(resource Resource) {
	key := resource.Key()

	rs.exists[key] = true

	_, exists := rs.resources[key]

	if exists == false {
		rs.resources[key] = make([]Resource, 0)
	}

	rs.resources[key] = append(rs.resources[key], resource)
}

func (rs ResourceSet) Remove(resource Resource) {
	key := resource.Key()

	delete(rs.exists, key)
	delete(rs.resources, key)
}

func (rs ResourceSet) Contains(resource Resource) bool {
	key := resource.Key()

	_, exists := rs.exists[key]

	return exists
}

func (rs ResourceSet) Resources() []Resource {
	resources := make([]Resource, 0)

	for _, slice := range rs.resources {
		for _, resource := range slice {
			resources = append(resources, resource)
		}
	}

	return resources
}

func (rs ResourceSet) Difference(other ResourceSet) {
	for _, resource := range rs.Resources() {
		if other.Contains(resource) {
			rs.Remove(resource)
		}
	}
}

func (rs ResourceSet) Unique() {
	for key, resources := range rs.resources {
		if len(resources) == 1 {
			continue
		}

		best := pick(resources)

		rs.resources[key] = []Resource{best}
	}
}

func pick(resources []Resource) Resource {
	best := resources[0]

	for _, resource := range resources[1:] {
		if resource.Fields() > best.Fields() {
			best = resource
		}
	}

	return best
}
