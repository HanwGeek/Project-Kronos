package geom

type Point struct {
	pos Coord
}

func (p Point) GeomType() int {
	return KrPoint
}

func (p Point) GetPos() (float32, float32) {
	return p.pos.x, p.pos.y
}

func (p *Point) SetPos(_x float32, _y float32) {
	p.pos.x = _x
	p.pos.y = _y
}
