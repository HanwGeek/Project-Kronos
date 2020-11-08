package geom

import "fmt"

type LineString struct {
	pos []Coord
}

func (l LineString) GeomType() int {
	return KrLineString
}

func (l LineString) NPoints() int {
	return len(l.pos)
}

func (l LineString) GetPoint(n int) Coord {
	m := l.NPoints()
	if n > m-1 {
		return l.pos[m-1] // Return the last point if out of the bound
	}
	return l.pos[n]
}

// Change the point with index n
func (l *LineString) EditPoint(n int, _x float64, _y float64) {
	m := l.NPoints()
	if n > m-1 {
		return
	}
	l.pos[n] = Coord{_x, _y}
}

// Add a point at the end
func (l *LineString) AddPoint(_x float64, _y float64) {
	l.pos = append(l.pos, Coord{_x, _y})
}

// Insert a point at index n
func (l *LineString) InsertPoint(n int, _x float64, _y float64) {
	m := l.NPoints()
	if n > m {
		return
	}
	rear := l.pos[n:]
	l.pos = append(l.pos[:n], Coord{_x, _y})
	l.pos = append(l.pos, rear...)
}

// Delete a point at index n
func (l *LineString) DeletePoint(n int) {
	m := l.NPoints()
	if n > m-1 {
		return
	}
	l.pos = append(l.pos[:n], l.pos[n+1:]...)
}

func (l LineString) ExportWKT() string {
	wkt := "LINESTRING ("
	for i := 0; i < len(l.pos); i++ {
		if i > 0 {
			wkt += ","
		}
		wkt += fmt.Sprintf("%f", l.pos[i].x) + " " + fmt.Sprintf("%f", l.pos[i].y)
	}
	wkt += ")"
	return wkt
}
