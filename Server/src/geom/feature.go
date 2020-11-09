package geom

// Feature means a combination of geometry and attributes.
type Feature struct {
	id   int
	geom Geometry
	attr map[string]interface{}
}

func NewFeatureByGeom(geom_ Geometry) *Feature {
	return &Feature{-1, geom_, make(map[string]interface{})}
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
