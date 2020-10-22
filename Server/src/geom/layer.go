package geom

// Layer means a collection of geoms
type Layer struct {
	geomtype int
	feat     []Feature
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

func (l Layer) GetFeature(n int) Feature {
	m := len(l.feat)
	if n > m-1 {
		return l.feat[m-1]
	}
	return l.feat[n]
}

// Add a feature
func (l *Layer) AddFeature(f Feature) {
	if l.geomtype != f.geom.GeomType() {
		return
	}
	l.feat = append(l.feat, f)
}

// Replace a feature at index n
func (l *Layer) ReplaceFeature(n int, f Feature) {
	if l.geomtype != f.geom.GeomType() {
		return
	}
	m := len(l.feat)
	if n > m-1 {
		return
	}
	l.feat[n]=f
}

// Delete a feature at index n
func (l *Layer) DeleteFeature(n int) {
	m := len(l.feat)
	if n > m-1 {
		return
	}
	l.feat = append(l.feat[:n], l.feat[n+1:]...)
}
