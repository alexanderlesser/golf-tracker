package models

// Club represents the club entity in the database
type Club struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Loft string `json:"loft"`
}
