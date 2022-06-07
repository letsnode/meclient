package types

import "time"

type LaunchpadCollections []struct {
	Symbol         string    `json:"symbol"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Featured       bool      `json:"featured"`
	Edition        string    `json:"edition"`
	Image          string    `json:"image"`
	Price          float64   `json:"price"`
	Size           int       `json:"size"`
	LaunchDatetime time.Time `json:"launchDatetime"`
}
