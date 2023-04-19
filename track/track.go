package track

import (
	"github.com/haydenheroux/cleanstring"
)

// Track represents a track.
type Track struct {
	// URL is the URL required to locate the track for download.
	URL string
	// Artists is a list of artists or creators of the track.
	Artists []string
	// Title is the title of the track.
	Title string
}

// String returns the representation of this track as a string.
func (t Track) String() string {
	artists := cleanstring.CleanSlice(t.Artists)
	title := cleanstring.Clean(t.Title)
	temp := map[string]string{artists: title}
	return cleanstring.CleanMap(temp)
}
