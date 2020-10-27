package ogcservice

import (
	"fmt"
	"kronos/src/geom"

	"github.com/go-pg/pg"
)

const (
	hostname = "162.105.17.227:5432"
	user     = "han"
	password = "wh123"
	dbname   = "features"
)

var db *pg.DB

// Manager manages the layers
type Manager struct {
	layers []geom.Layer
}

type LayerInfo struct {
	tableName struct{} `pg:"layer_info"`
	LayerID   int      `pg:"layer_id"`
	LayerName string   `pg:"layer_name"`
	Type      int      `pg:"type"`
	Count     int      `pg:"count"`
}

// Connect is to connect postgresql database
func Connect() {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     hostname,
			User:     user,
			Password: password,
			Database: dbname,
		})
	}
	GetLayerInfo()
}

// GetLayerInfo return layer metadata
func GetLayerInfo() []LayerInfo {
	var model []LayerInfo
	err := db.Model(&model).Select()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil
	}
	return model
}
