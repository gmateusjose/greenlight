package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // not show CreatedAt field
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`           // show if field is non-zero
	Runtime   int32     `json:"runtime,omitempty,string"` // show if field is non-zero
	Genres    []string  `json:"genres,omitempty"`         // show if field is non-zero
	Version   int32     `json:"version"`
}
