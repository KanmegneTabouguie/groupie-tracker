// date.go

package models

type Date struct {
	ID       int      `json:"id"`
	Dates    []string `json:"dates"`
	ArtistID int      `json:"artistId"` // Add this line

}
