package entity

type Artist struct {
	ID int    `db:"ArtistId"`
	Name      string `db:"Name"`
}
