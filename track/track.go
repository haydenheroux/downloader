package track

// Track represents a track.
type Track struct {
	// URL is the URL required to locate the track for download.
	URL string
	// Name is the name of the track.
	Name string
}

// String returns the representation of this track as a string.
func (t Track) String() string {
	return t.Name
}
