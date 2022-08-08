package model

type Song struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	ReleaseDate string `json:"releaseDate"`
}
