package resource

import (
	"strings"

	"github.com/haydenheroux/strfmt"
)

// namedUrl represents a URL with an associated name.
type namedUrl struct {
	// url is the URL that is being named.
	url string
	// name is the name associated with the URL.
	name string
}

func (n namedUrl) PrimaryKey() primaryKey {
	return primaryKey(n.name)
}

func (n namedUrl) Title() string {
	return n.name
}

func (n namedUrl) Source() string {
	return n.url
}

func (n namedUrl) MetadataFields() int {
	// url and name
	return 2
}

// attributedUrl represents a URL with associated creators and a descriptive title.
type attributedUrl struct {
	// url is the URL that is being represented.
	url string
	// creators is a slice containing the creators of the content.
	creators []string
	// name is the name of the content.
	name string
}

func createAttributedUrl(fields []string) attributedUrl {
	return attributedUrl{
		url:      fields[0],
		creators: strings.Split(fields[1], "&"),
		name:     fields[2],
	}
}

func (a attributedUrl) PrimaryKey() primaryKey {
	return primaryKey(a.Title())
}

func (a attributedUrl) Title() string {
	creators := strfmt.Join(a.creators)

	return strfmt.Associate(map[string]string{creators: a.name})
}

func (a attributedUrl) Source() string {
	return a.url
}

func (a attributedUrl) MetadataFields() int {
	// url, creators, name
	return 3
}

// taggedResource adds tags to an existing resource.
type taggedResource struct {
	resource Resource
	tags     []string
}

func (tr taggedResource) PrimaryKey() primaryKey {
	// Use the primary key of the nested resource
	return tr.resource.PrimaryKey()
}

func (tr taggedResource) Title() string {
	title := tr.resource.Title()

	tags := strfmt.Join(tr.tags)

	return strfmt.Join([]string{title, tags})
}

func (tr taggedResource) Source() string {
	return tr.resource.Source()
}

func (tr taggedResource) MetadataFields() int {
	return tr.resource.MetadataFields() + len(tr.tags)
}
