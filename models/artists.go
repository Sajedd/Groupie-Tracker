package model

type Artist struct {
	Id           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	LocationsUrl string `json:"locations"`
	//Locations []Location

	ConcertDatesUrl string `json:"concertDates"`
	// ConcertDates []string `json:"Dates"`

	RelationsUrl string `json:"relations"`
	//Relations []Relation
}

var (
	Artists Artist
)
