package geom

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
	f.SetID(l.next_id)
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
