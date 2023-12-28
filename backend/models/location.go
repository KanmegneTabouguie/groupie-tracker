// location.go

package models

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
	ArtistID  int      `json:"artistId"` // Add this line

}
