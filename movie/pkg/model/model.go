package model

import "movieexample.com/metadata/pkg/model"

// MovieDetails includes novie metadata its aggregated rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}