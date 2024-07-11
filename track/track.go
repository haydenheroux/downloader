package track

import (
	fmt "github.com/haydenheroux/strfmt"
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
	artists := fmt.Join(t.Artists)
	title := fmt.Format(t.Title)
    return fmt.Associate(map[string]string{artists: title})
}
