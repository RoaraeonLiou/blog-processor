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
}
