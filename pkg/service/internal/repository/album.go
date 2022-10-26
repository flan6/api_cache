package repository

import (
	"database/sql"
	"fmt"

	"github.com/ReneKroon/ttlcache"

	"api/pkg/entity"
)

type Album interface {
	Get(id int) (*entity.Album, error)
}

type albumRepository struct {
	db    *sql.DB
	cache *ttlcache.Cache
}

func NewAlbum(db *sql.DB, cache *ttlcache.Cache) Album {
	return albumRepository{
		db:    db,
		cache: cache,
	}
}

func albumCacheKey(id int) string {
	return fmt.Sprintf("%T:%d", entity.Album{}, id)
}

func (a albumRepository) Get(id int) (*entity.Album, error) {
	album, exists := a.cache.Get(albumCacheKey(id))
	if !exists {
		var album entity.Album
		query := fmt.Sprintf("select AlbumId, Title, ArtistId from albums where AlbumId = %d", id)

		row := a.db.QueryRow(query)

		err := row.Scan(&album.ID, &album.Title, &album.ArtistID)
		if err != nil {
			return nil, err
		}

		a.cache.Set(albumCacheKey(id), album)
		return &album, nil
	}

	if a, ok := album.(entity.Album); ok {
		return &a, nil
	}

	return nil, nil
}
