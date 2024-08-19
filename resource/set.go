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
	for key, slice := range rs.resources {
		if len(slice) == 1 {
			continue
		}

		rs.resources[key] = unique(rs.resources[key])
	}
}

func unique(resources []Resource) []Resource {
	simplified := make([]Resource, 0)

	byName := groupByName(resources)

	for _, resources := range byName {
		picked := pick(resources)

		simplified = append(simplified, picked)
	}

	return simplified
}

func groupByName(resources []Resource) map[string][]Resource {
	byName := make(map[string][]Resource)

	for _, resource := range resources {
		name := resource.Name()

		_, ok := byName[name]

		if ok == false {
			byName[name] = make([]Resource, 0)
		}

		byName[name] = append(byName[name], resource)
	}

	return byName
}

func pick(resources []Resource) Resource {
	return resources[0]
}
