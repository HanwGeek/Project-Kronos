package ogcservice

import (
	"kronos/src/geom"
)

// LayerManager manage layers in memory
type LayerManager struct {
	layers   map[int]*geom.Layer
	layerIDs []int
}

func NewManager() *LayerManager {
	return &LayerManager{
		layers: make(map[int]*geom.Layer),
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
	lm.layers[layerID] = &layer

	return layer.ExportMap()
}

// OperOnLayer operates `ops` on layer for a `POST` request
func (lm *LayerManager) OperOnLayer(OpID int, LayerID int) {
	switch OpID {
	case 0:

	}
}

func (lm *LayerManager) AddFeatureToLayer(LayerID int, feats map[string]interface{}) {
	for _, feat := range feats {
		lm.layers[LayerID].AddFeature(*geom.NewFeatureFromJSON(feat.(map[string]interface{})))
	}
}

func (lm *LayerManager) EditFeatureToLaye(LayerID int, feats map[string]interface{}) {

}
