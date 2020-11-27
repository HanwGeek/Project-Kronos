package geom

import (
	"encoding/json"
)

// Feature means a combination of geometry and attributes.
type Feature struct {
	id   int
	geom Geometry
	attr map[string]interface{}
}

func NewFeatureByGeom(geom_ Geometry) *Feature {
	return &Feature{-1, geom_, make(map[string]interface{})}
}

func NewFeatureFromJSON(json_ map[string]interface{}) *Feature {
	// var m map[string]interface{}
	// err := json.Unmarshal([]byte(json_), &m)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return nil
	// }

	m := json_
	var feat_temp Feature
	switch geom_ := m["geometry"].(type) {
	case map[string]interface{}:
		if geom_["type"] == "Point" {
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
	switch attr_ := m["properties"].(type) {
	case map[string]interface{}:
		feat_temp.attr = attr_
	}
	return &feat_temp
}

// Add by Ganmin Yin.
func NewFeature(geom_ Geometry, attr_ map[string]interface{}) *Feature {
	return &Feature{-1, geom_, attr_}
}

func (feat Feature) GetID() int {
	return feat.id
}

func (feat *Feature) SetID(id_ int) {
	feat.id = id_
}

// The function to get the feature's type.
func (feat Feature) GetGeomType() int {
	return feat.geom.GeomType()
}

// The function to get the feature's geometry.
func (feat Feature) GetGeometry() Geometry {
	return feat.geom
}

func (feat *Feature) SetGeometry(g Geometry) {
	if g.GeomType() != feat.geom.GeomType() {
		return
	}
	feat.geom = g
}

// The function to get the feature's attributes.
func (feat Feature) GetAttribute(attrname string) interface{} {
	return feat.attr[attrname]
}

// The function to get the feature's attributes.
func (feat Feature) GetAttributes() map[string]interface{} {
	return feat.attr
}

func (feat *Feature) SetAttribute(attrname string, val interface{}) {
	feat.attr[attrname] = val
}

func (feat *Feature) DeleteAttribute(attrname string) {
	_, ok := feat.attr[attrname]
	if ok {
		delete(feat.attr, attrname)
	}
}

func (feat *Feature) ExportMap() map[string]interface{} {
	mj := make(map[string]interface{})
	mj["type"] = "Feature"
	mj["properties"] = feat.attr
	mj["geometry"] = feat.geom.ExportMap()
	return mj
}

func (feat *Feature) ExportGeoJSON() string {
	s, _ := json.Marshal(feat.ExportMap())
	return string(s)
}
