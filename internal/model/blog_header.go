package model

import (
	"blog-processor/pkg/utils/set_operation"
	"reflect"
)

type BlogHeader struct {
	Title             string    `yaml:"title" toml:"title" json:"title"`
	Description       string    `yaml:"description,omitempty" toml:"description,omitempty" json:"description,omitempty"`
	Tags              []string  `yaml:"tags,omitempty" toml:"tags,omitempty" json:"tags,omitempty"`
	Date              string    `yaml:"date,omitempty" toml:"date,omitempty" json:"date,omitempty"`
	Categories        []string  `yaml:"categories,omitempty" toml:"categories,omitempty" json:"categories,omitempty"`
	Slug              string    `yaml:"slug,omitempty" toml:"slug,omitempty" json:"slug,omitempty"`
	Author            string    `yaml:"author,omitempty" toml:"author,omitempty" json:"author,omitempty"`
	Draft             bool      `yaml:"draft" toml:"draft" json:"draft"`
	Cover             BlogCover `yaml:"cover,omitempty" toml:"cover,omitempty" json:"cover,omitempty"`
	Status            string    `yaml:"status,omitempty" toml:"status,omitempty" json:"status,omitempty"`
	LastMod           string    `yaml:"lastmod,omitempty" toml:"lastmod,omitempty" json:"lastmod,omitempty"`
	Layout            string    `yaml:"layout,omitempty" toml:"layout,omitempty" json:"layout,omitempty"`
	Summary           string    `yaml:"summary,omitempty" toml:"summary,omitempty" json:"summary,omitempty"`
	PlaceHolder       string    `yaml:"placeholder,omitempty" toml:"placeholder,omitempty" json:"placeholder,omitempty"`
	Url               string    `yaml:"url,omitempty" toml:"url,omitempty" json:"url,omitempty"`
	Type              string    `yaml:"type,omitempty" toml:"type,omitempty" json:"type,omitempty"`
	IncludeTags       []string  `yaml:"include_tags,omitempty" toml:"include_tags,omitempty" json:"include_tags,omitempty"`
	ExcludeTags       []string  `yaml:"exclude_tags,omitempty" toml:"exclude_tags,omitempty" json:"exclude_tags,omitempty"`
	IncludeCategories []string  `yaml:"include_categories,omitempty" toml:"include_categories,omitempty" json:"include_categories,omitempty"`
	ExcludeCategories []string  `yaml:"exclude_categories,omitempty" toml:"exclude_categories,omitempty" json:"exclude_categories,omitempty"`
}

func (header *BlogHeader) Merge(other *BlogHeader) {
	headerValue := reflect.ValueOf(header).Elem()
	otherValue := reflect.ValueOf(other).Elem()

	for i := 0; i < headerValue.NumField(); i++ {
		headerField := headerValue.Field(i)
		otherField := otherValue.Field(i)

		if isZeroValue(headerField) && !isZeroValue(otherField) {
			headerField.Set(otherField)
		}
	}

	header.Tags = set_operation.Union(header.Tags, header.IncludeTags)
	header.Tags = set_operation.Exclude(header.Tags, header.ExcludeTags)
	header.Categories = set_operation.Union(header.Categories, header.IncludeCategories)
	header.Categories = set_operation.Exclude(header.Categories, header.ExcludeCategories)
}

func isZeroValue(v reflect.Value) bool {
	zeroValue := reflect.Zero(v.Type()).Interface()
	return reflect.DeepEqual(v.Interface(), zeroValue)
}
