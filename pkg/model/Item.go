package model

import "time"

type Items []Item

type Item struct {
	ID string

	Metrics map[string]uint

	Created time.Time
}
