package model

import "time"

type Song struct {
	id          int       `json:"id"`
	name        string    `json:"name"`
	author      string    `json:"author"`
	releaseDate time.Time `json:"releaseDate"`
}
