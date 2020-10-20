package ogcservice

import (
	"kronos/src/geom"

	"github.com/go-pg/pg"
)

const (
	hostname = "162.105.17.227:5432"
	user     = "han"
	password = "wh123"
	dbname   = "mapdb"
)

var db *pg.DB

// Manager manages the layers
type Manager struct {
	layers []geom.Layer
}

func connect() {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     hostname,
			User:     user,
			Password: password,
			Database: dbname,
		})
	}
}
