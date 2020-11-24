package geom

// Coord is the basic type of gemotry with 2-D position
type Coord struct {
	x float64
	y float64
}

// Coord is the basic type of gemotry with 3-D position
type Coord3D struct {
	x float64
	y float64
	z float64
}

func NewCoord(pos []float64) Coord {
	return Coord{pos[0], pos[1]}
}

func NewCoordList(pos [][]float64) []Coord {
	var coordList []Coord
	for _, c := range pos {
		coordList = append(coordList, NewCoord(c))
	}
	return coordList
}
