package service

import (
	"database/sql"

	"github.com/ReneKroon/ttlcache"

	"api/pkg/entity"
	"api/pkg/service/internal/repository"
)

type Artist interface {
	Get(id int) (*entity.Artist, error)
}

func NewArtist(db *sql.DB, cache *ttlcache.Cache) Artist {
	return artist{
		artistRepository: repository.NewArtist(db, cache),
	}
}

type artist struct {
	artistRepository repository.Artist
}

func (a artist) Get(id int) (*entity.Artist, error) {
	return a.artistRepository.Get(id)
}
