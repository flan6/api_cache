package service

import (
	"database/sql"

	"github.com/ReneKroon/ttlcache"
)

type All struct {
	Album  Album
	Artist Artist
}

func GetAll(db *sql.DB, cache *ttlcache.Cache) All {
	return All{
		Album:  NewAlbum(db, cache),
		Artist: NewArtist(db, cache),
	}
}
