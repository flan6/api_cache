package repository

import (
	"database/sql"
	"fmt"

	"github.com/ReneKroon/ttlcache"

	"api/pkg/entity"
)

type Artist interface {
	Get(id int) (*entity.Artist, error)
}

type artistRepository struct {
	db    *sql.DB
	cache *ttlcache.Cache
}

func NewArtist(db *sql.DB, cache *ttlcache.Cache) Artist {
	return artistRepository{
		db:    db,
		cache: cache,
	}
}

func artistCacheKey(id int) string {
	return fmt.Sprintf("%T:%d", entity.Artist{}, id)
}

func (a artistRepository) Get(id int) (*entity.Artist, error) {
	artist, exists := a.cache.Get(artistCacheKey(id))
	if !exists {
		var artist entity.Artist
		query := fmt.Sprintf("select ArtistId, Name from artists where ArtistId = %d", id)

		row := a.db.QueryRow(query)

		err := row.Scan(&artist.ID, &artist.Name)
		if err != nil {
			return nil, err
		}

		a.cache.Set(artistCacheKey(id), artist)
		return &artist, nil
	}

	if a, ok := artist.(entity.Artist); ok {
		return &a, nil
	}

	return nil, nil
}
