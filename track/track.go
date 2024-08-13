package track

import (
	"github.com/haydenheroux/strfmt"
)

type Track interface {
	URL() string
	Name() string
}

type urlName struct {
	url  string
	name string
}

func (t urlName) URL() string {
	return t.url
}

func (t urlName) Name() string {
	return t.name
}

type urlArtistsTitle struct {
	url     string
	artists []string
	title   string
}

func (t urlArtistsTitle) URL() string {
	return t.url
}

func (t urlArtistsTitle) Name() string {
	artists := strfmt.Join(t.artists)

	return strfmt.Associate(map[string]string{artists: t.title})
}
