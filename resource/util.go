package resource

func Difference(resources []Resource, set map[Resource]bool) []Resource {
	result := make([]Resource, 0, len(resources))

	for _, track := range resources {
		if exists, _ := set[track]; !exists {
			result = append(result, track)
		}
	}

	return result
}

func Unique(tracks []Resource) []Resource {
	result := make([]Resource, 0, len(tracks))

	set := make(map[string]bool)

	for _, track := range tracks {
		name := track.Name()

		if exists, _ := set[name]; !exists {
			result = append(result, track)
			set[name] = true
		}
	}

	return result
}
