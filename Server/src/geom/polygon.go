package geom

import "fmt"

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

func (p Polygon) ExportWKT() string {
	wkt := "POLYGON ("
	for i := 0; i < len(p.rings); i++ {
		if i > 0 {
			wkt += ","
		}
		wkt += "("
		for j := 0; j < len(p.rings[i].pos); j++ {
			if j > 0 {
				wkt += ","
			}
			wkt += fmt.Sprintf("%f", p.rings[i].pos[j].x) + " " + fmt.Sprintf("%f", p.rings[i].pos[j].y)

		}
		wkt += "," + fmt.Sprintf("%f", p.rings[i].pos[0].x) + " " + fmt.Sprintf("%f", p.rings[i].pos[0].y)
		wkt += ")"
	}
	wkt += ")"
	return wkt
}
