package geom

// Coord is the basic type of gemotry with 3-D position
type Coord struct {
	x float32
	y float32
	z float32
}

func (c *Coord) isNull() {
}
