package resource

import (
	"strings"

	"github.com/haydenheroux/strfmt"
)

type key string

type Resource interface {
	Key() key
	Name() string
	Source() string
	Specificity() int
}

type namedUrl struct {
	url  string
	name string
}

func (n namedUrl) Key() key {
	return key(n.Name())
}

func (n namedUrl) Name() string {
	return n.name
}

func (n namedUrl) Source() string {
	return n.url
}

func (n namedUrl) Specificity() int {
	return 2
}

type attributedUrl struct {
	url     string
	artists []string
	title   string
}

func createAttributedURL(fields []string) attributedUrl {
	return attributedUrl{
		url:     fields[0],
		artists: strings.Split(fields[1], "&"),
		title:   fields[2],
	}
}

func (a attributedUrl) Key() key {
	return key(a.Name())
}

func (a attributedUrl) Name() string {
	artists := strfmt.Join(a.artists)

	return strfmt.Associate(map[string]string{artists: a.title})
}

func (a attributedUrl) Source() string {
	return a.url
}

func (a attributedUrl) Specificity() int {
	return 3
}

type taggedResource struct {
	resource Resource
	tags     []string
}

func (tr taggedResource) Key() key {
	return tr.resource.Key()
}

func (tr taggedResource) Name() string {
	name := tr.resource.Name()

	tags := strfmt.Join(tr.tags)

	return strfmt.Join([]string{name, tags})
}

func (tr taggedResource) Source() string {
	return tr.resource.Source()
}

func (tr taggedResource) Specificity() int {
	return tr.resource.Specificity() + len(tr.tags)
}
