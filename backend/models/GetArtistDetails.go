// artist_details.go
package models

// ArtistDetails represents detailed information about an artist, including locations, dates, and relations
type ArtistDetails struct {
	Artist    Artist        `json:"artist"`
	Locations []Location    `json:"locations"`
	Dates     []interface{} `json:"dates"`
	Relations []interface{} `json:"relations"`
}
