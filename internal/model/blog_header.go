package model

type BlogHeader struct {
	Title       string    `yaml:"title" toml:"title" json:"title"`
	Description string    `yaml:"description,omitempty" toml:"description,omitempty" json:"description,omitempty"`
	Tags        []string  `yaml:"tags,omitempty" toml:"tags,omitempty" json:"tags,omitempty"`
	Date        string    `yaml:"date,omitempty" toml:"date,omitempty" json:"date,omitempty"`
	Categories  []string  `yaml:"categories,omitempty" toml:"categories,omitempty" json:"categories,omitempty"`
	Slug        string    `yaml:"slug,omitempty" toml:"slug,omitempty" json:"slug,omitempty"`
	Author      string    `yaml:"author,omitempty" toml:"author,omitempty" json:"author,omitempty"`
	Draft       bool      `yaml:"draft" toml:"draft" json:"draft"`
	Cover       BlogCover `yaml:"cover,omitempty" toml:"cover,omitempty" json:"cover,omitempty"`
	Status      string    `yaml:"-" toml:"-" json:"-"`
	LastMod     string    `yaml:"lastmod,omitempty" toml:"lastmod,omitempty" json:"lastmod,omitempty"`
	Layout      string    `yaml:"layout,omitempty" toml:"layout,omitempty" json:"layout,omitempty"`
	Summary     string    `yaml:"summary,omitempty" toml:"summary,omitempty" json:"summary,omitempty"`
	PlaceHolder string    `yaml:"placeholder,omitempty" toml:"placeholder,omitempty" json:"placeholder,omitempty"`
	Url         string    `yaml:"url,omitempty" toml:"url,omitempty" json:"url,omitempty"`
}
