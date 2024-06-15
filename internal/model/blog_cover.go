package model

type BlogCover struct {
	Image    string `yaml:"image,omitempty"`
	Alt      string `yaml:"alt,omitempty"`
	Caption  string `yaml:"caption,omitempty"`
	Relative bool   `yaml:"relative,omitempty"`
	Hidden   bool   `yaml:"hidden,omitempty"`
}
