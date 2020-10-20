package geom

type Polygon struct {
	// rings[0] is the outer ring, the others are the inner rings
	// In each LineString, the last point is linked to the first point
	rings []LineString
}

func (p Polygon) GeomType() int {
	return KrPolygon
}

func (p Polygon) NInnerRings() int {
	return len(p.rings) - 1
}

func (p Polygon) GetRing(n int) LineString {
	m := len(p.rings)
	if n > m-1 {
		return p.rings[m-1]
	}
	return p.rings[n]
}

func (p *Polygon) AddRing(l LineString) {
	p.rings = append(p.rings, l)
}

// Change the point at index n
func (p *Polygon) EditRing(n int, l LineString) {
	m := len(p.rings)
	if n > m-1 {
		return
	}
	p.rings[n] = l
}

func (p *Polygon) DeleteRing(n int) {
	m := len(p.rings)
	// Delete the Outer Ring is not allowed
	if n > m-1 || n <= 0 {
		return
	}
	p.rings = append(p.rings[:n], p.rings[n+1:]...)
}
