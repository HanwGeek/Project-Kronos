package geom

// Feature means a combination of geometry and attributes.
type Feature struct {
	geom Geometry
	attr map[string]interface{}
}

// The function to get the feature's type.
func (feat Feature) GetGeomType() int {
	return feat.geom.GeomType()
}

// The function to get the feature's geometry.
func (feat Feature) GetGeometry() Geometry {
	return feat.geom
}

// The function to get the feature's attributes.
func (feat Feature) GetAttributes() map[string]interface{} {
	return feat.attr
}

// The function waiting to discuss.
func (feat Feature) EditFeature() {

}
