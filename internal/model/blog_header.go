package model

type BlogHeader struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description,omitempty"`
	Tags        []string  `yaml:"tags,omitempty"`
	Date        string    `yaml:"date,omitempty"`
	Categories  []string  `yaml:"categories,omitempty"`
	Slug        string    `yaml:"slug,omitempty"`
	Author      string    `yaml:"author,omitempty"`
	Draft       bool      `yaml:"draft,omitempty"`
	Cover       BlogCover `yaml:"cover,omitempty"`
	Status      string    `yaml:"-"`
	LastMod     string    `yaml:"lastmod,omitempty"`
	Layout      string    `yaml:"layout,omitempty"`
	Summary     string    `yaml:"summary,omitempty"`
	PlaceHolder string    `yaml:"placeholder,omitempty"`
	Url         string    `yaml:"url,omitempty"`
}
