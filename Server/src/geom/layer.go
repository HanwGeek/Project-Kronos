package geom

// Layer means a collection of geoms
type Layer struct {
	geomtype int
	feat []Geometry
}

func (l Layer) GetGeomType () int {
	return l.geomtype
}

func (l Layer) IsEmpty () bool {
	return len(l.feat)==0
}

func (l Layer) FeatureCount () int {
	return len(l.feat)
}

func (l Layer) GetFeature (n int) Geometry {
	m := len(l.feat)
	if n > m-1 {
		return l.feat[m-1]
	}
	return l.feat[n]
}

// Add a feature at the end
func (l *Layer) AddFeature(g Geometry) {
	if l.geomtype!=g.GeomType() {
		return
	}
	l.feat=append(l.feat,g)
}

// Delete a point at index n
func (l *Layer) DeleteFeature(n int) {
	m := len(l.feat)
	if n > m - 1 {
		return
	}
	l.feat = append(l.feat[:n], l.feat[n+1:]...)
}
