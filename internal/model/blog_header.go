package model

type BlogHeader struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Date        string    `json:"date"`
	Categories  []string  `json:"categories"`
	Slug        string    `json:"slug"`
	Author      string    `json:"author"`
	Draft       bool      `json:"draft"`
	Cover       BlogCover `json:"cover"`
}
