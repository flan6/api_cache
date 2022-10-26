package service

import (
	"database/sql"

	"github.com/ReneKroon/ttlcache"

	"api/pkg/entity"
	"api/pkg/service/internal/repository"
)

type Album interface {
	Get(id int) (*entity.Album, error)
}

func NewAlbum(db *sql.DB, cache *ttlcache.Cache) Album {
	return album{
		albumRepository: repository.NewAlbum(db, cache),
	}
}

type album struct {
	albumRepository repository.Album
}

func (a album) Get(id int) (*entity.Album, error) {
	return a.albumRepository.Get(id)
}
