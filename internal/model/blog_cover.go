package model

type BlogCover struct {
	Image    string `json:"image"`
	Alt      string `json:"alt"`
	Caption  string `json:"caption"`
	Relative bool   `json:"relative"`
	Hidden   bool   `json:"hidden"`
}
