package entity

type Album struct {
	ID       int    `db:"AlbumId"`
	Title    string `db:"Title"`
	ArtistID int    `db:"ArtistId"`
}
