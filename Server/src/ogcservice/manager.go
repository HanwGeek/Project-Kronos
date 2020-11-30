package ogcservice

import (
	"kronos/src/geom"
)

// LayerManager manage layers in memory
type LayerManager struct {
	layers map[int]*geom.Layer
	cache  map[int]*LayerCache
	queue  []int
	cap    int
}

type LayerCache struct {
	add    []*geom.Feature
	edit   []*geom.Feature
	delete []int
}

func NewManager() *LayerManager {
	return &LayerManager{
		layers: make(map[int]*geom.Layer),
		cache:  make(map[int]*LayerCache),
		cap:    5,
	}
}

// GetLayerinfo return layer info for a `GET` request
func (lm *LayerManager) GetLayerinfo() []LayerInfo {
	return GetLayerInfo()
}

// GetLayerContent return layer content for a `GET` request
func (lm *LayerManager) GetLayerContent(layerID int) map[string]interface{} {
	// if layer is loaded
	if layer, ok := lm.layers[layerID]; ok {
		return layer.ExportMap()
	}

	// Load layer
	layer := GetLayerById(layerID)

	if len(lm.layers) >= lm.cap {
		// Dump layer
		dump := lm.queue[0]
		lm.DumpToDatabase(dump)
		delete(lm.layers, dump)
		delete(lm.cache, dump)
		lm.queue = lm.queue[1:]
		lm.cap--
	}

	lm.layers[layerID] = &layer
	lm.queue = append(lm.queue, layerID)
	lm.cap++
	return layer.ExportMap()
}

// OperOnLayer operates `ops` on layer for a `POST` request
func (lm *LayerManager) OperOnLayer(OpID int, LayerID int, Feat map[string]interface{}) {
	if _, ok := lm.layers[LayerID]; !ok {
		// Load layer
		layer := GetLayerById(LayerID)

		if len(lm.layers) >= lm.cap {
			// Dump layer
			dump := lm.queue[0]
			lm.DumpToDatabase(dump)
			delete(lm.layers, dump)
			delete(lm.cache, dump)
			lm.queue = lm.queue[1:]
			lm.cap--
		}

		lm.layers[LayerID] = &layer
		lm.queue = append(lm.queue, LayerID)
		lm.cap++
	}
	feats := Feat["features"]
	switch OpID {
	case 1:
		lm.AddFeatureToLayer(LayerID, feats.([]interface{}))
	case 2:
		lm.EditFeatureToLayer(LayerID, feats.([]interface{}))
	}
}

func (lm *LayerManager) AddFeatureToLayer(LayerID int, feats []interface{}) {
	for _, feat := range feats {
		newFeat := geom.NewFeatureFromJSON(feat.(map[string]interface{}))
		lm.layers[LayerID].AddFeature(*newFeat)
		lm.cache[LayerID].add = append(lm.cache[LayerID].add, newFeat)
	}
}

func (lm *LayerManager) EditFeatureToLayer(LayerID int, feats []interface{}) {
	for _, feat := range feats {
		newFeat := geom.NewFeatureFromJSON(feat.(map[string]interface{}))
		lm.layers[LayerID].ReplaceFeature(newFeat.GetID(), *newFeat)
		lm.cache[LayerID].edit = append(lm.cache[LayerID].edit, newFeat)
	}
}

func (lm *LayerManager) DumpToDatabase(LayerID int) {
	for _, feat := range lm.cache[LayerID].add {
		AddFeature(LayerID, *feat)
	}

	for _, feat := range lm.cache[LayerID].edit {
		UpdateFeature(LayerID, *feat)
	}

	for _, feat := range lm.cache[LayerID].delete {
		DelFeature(LayerID, feat)
	}
}
