package service

import "database/sql"

type All struct {
	Album Album
}

func GetAll(db *sql.DB) All {
	return All{
		Album: NewAlbum(db),
	}
}
