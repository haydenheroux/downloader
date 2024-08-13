package resource

import (
	"github.com/haydenheroux/strfmt"
)

type Resource interface {
	Name() string
	Source() string
}

type urlName struct {
	url  string
	name string
}

func (t urlName) Name() string {
	return t.name
}

func (t urlName) Source() string {
	return t.url
}

type urlArtistsTitle struct {
	url     string
	artists []string
	title   string
}

func (t urlArtistsTitle) Name() string {
	artists := strfmt.Join(t.artists)

	return strfmt.Associate(map[string]string{artists: t.title})
}

func (t urlArtistsTitle) Source() string {
	return t.url
}
