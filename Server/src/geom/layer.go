package geom

import (
	"encoding/json"
	"fmt"
)

// Layer means a collection of geoms
type Layer struct {
	id       int // Id for layer, as the only identifier
	name     string
	geomtype int
	feat     map[int]Feature
	// Id for next object. All features must be added by AddFeature()
	next_id int
}

// Add id.
func NewLayer(id_ int, name_ string, geomtype_ int) *Layer {
	return &Layer{id_, name_, geomtype_, make(map[int]Feature), 0}
}

func NewLayerfromGeoJSON(json_ string, id_ int) *Layer {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(json_), &m)
	if err != nil {
		fmt.Printf("err = %v", err)
		return nil
	}

	var layer Layer
	layer.id = id_
	layer.feat = make(map[int]Feature)
	switch featlist := m["features"].(type) {
	case []interface{}:
		for i := 0; i < len(featlist); i++ {
			var feat_temp Feature
			switch feat_ := featlist[i].(type) {
			case map[string]interface{}:
				switch geom_ := feat_["geometry"].(type) {
				case map[string]interface{}:
					if geom_["type"] == "Point" {
						if i == 0 {
							layer.geomtype = KrPoint
						}
						var geom_temp Point
						switch pos_ := geom_["coordinates"].(type) {
						case []interface{}:
							switch x_ := pos_[0].(type) {
							case float64:
								switch y_ := pos_[1].(type) {
								case float64:
									geom_temp.SetPos(x_, y_)
								}
							}
						}
						feat_temp = *NewFeatureByGeom(geom_temp)
					} else if geom_["type"] == "LineString" {
						if i == 0 {
							layer.geomtype = KrLineString
						}
						var geom_temp LineString
						switch pos_ := geom_["coordinates"].(type) {
						case []interface{}:
							for j := 0; j < len(pos_); j++ {
								switch point_ := pos_[j].(type) {
								case []interface{}:
									switch x_ := point_[0].(type) {
									case float64:
										switch y_ := point_[1].(type) {
										case float64:
											geom_temp.AddPoint(x_, y_)
										}
									}
								}
							}
						}
						feat_temp = *NewFeatureByGeom(geom_temp)
					} else {
						if i == 0 {
							layer.geomtype = KrPolygon
						}
						var geom_temp Polygon
						switch pos_ := geom_["coordinates"].(type) {
						case []interface{}:
							for r := 0; r < len(pos_); r++ {
								var ring_temp LineString
								switch ring_ := pos_[r].(type) {
								case []interface{}:
									for j := 0; j < len(ring_)-1; j++ {
										switch point_ := ring_[j].(type) {
										case []interface{}:
											switch x_ := point_[0].(type) {
											case float64:
												switch y_ := point_[1].(type) {
												case float64:
													ring_temp.AddPoint(x_, y_)
												}
											}
										}
									}
									geom_temp.AddRing(ring_temp)
								}
							}
						}
						feat_temp = *NewFeatureByGeom(geom_temp)
					}
				}
				switch attr_ := feat_["properties"].(type) {
				case map[string]interface{}:
					feat_temp.attr = attr_
				}
				layer.AddFeature(feat_temp)
			}
		}
	}
	return &layer
}

func (l Layer) GetId() int {
	return l.id
}

func (l Layer) GetName() string {
	return l.name
}

func (l *Layer) SetName(name_ string) {
	l.name = name_
}

func (l Layer) GetGeomType() int {
	return l.geomtype
}

func (l Layer) IsEmpty() bool {
	return len(l.feat) == 0
}

func (l Layer) FeatureCount() int {
	return len(l.feat)
}

func (l Layer) GetFeature(id_ int) Feature {
	feat, ok := l.feat[id_]
	if ok {
		return feat
	}
	return Feature{}
}

// Add a feature
func (l *Layer) AddFeature(f Feature) {
	if l.geomtype != f.geom.GeomType() {
		return
	}
	if f.GetID() == -1 {
		f.SetID(l.next_id)
	}
	l.feat[l.next_id] = f
	l.next_id += 1
}

// Replace a feature at index n
func (l *Layer) ReplaceFeature(id_ int, f Feature) {
	if l.geomtype != f.geom.GeomType() {
		return
	}
	_, ok := l.feat[id_]
	if !ok {
		return
	}
	l.feat[id_] = f
}

// Delete a feature at index n
func (l *Layer) DeleteFeature(id_ int) {
	_, ok := l.feat[id_]
	if !ok {
		return
	}
	delete(l.feat, id_)
}

func (l Layer) ExportMap() map[string]interface{} {
	mj := make(map[string]interface{})
	mj["type"] = "FeatureCollection"
	var featArray []map[string]interface{}
	for id := range l.feat {
		feature := l.feat[id]
		featArray = append(featArray, feature.ExportMap())
	}
	mj["features"] = featArray
	return mj
}

func (l Layer) ExportGeoJSON() string {
	mj := l.ExportMap()
	s, _ := json.Marshal(mj)
	return string(s)
}
