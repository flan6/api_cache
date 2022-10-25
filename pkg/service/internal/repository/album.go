package repository

import (
	"database/sql"
	"fmt"

	"api/pkg/entity"
)

type Album interface {
	Get(id int) (*entity.Album, error)
}

type albumRepository struct {
	db *sql.DB
}

func NewAlbum(db *sql.DB) Album {
	return albumRepository{
		db: db,
	}
}

func (a albumRepository) Get(id int) (*entity.Album, error) {
	var album entity.Album
	query := fmt.Sprintf("select AlbumId, Title, ArtistId from albums where AlbumId = %d", id)

	row := a.db.QueryRow(query)

	err := row.Scan(&album.ID, &album.Title, &album.ArtistID)
	if err != nil {
		return nil, err
	}

	return &album, nil
}
