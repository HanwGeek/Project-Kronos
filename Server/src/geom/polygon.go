package geom

import (
	"encoding/json"
	"fmt"
)

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

func (p Polygon) ExportMap() map[string] interface{} {
	mj:=make(map[string]interface{})
	mj["type"]="MultiPolygon"
	var geomArray [][][][] float64
	var polygonArray[][][] float64
	for i := 0; i < p.NInnerRings()+1; i++ {
		n:=p.rings[i].NPoints()
		var ringArray [][] float64
		var pointArray [] float64
		for j:= 0; j < n ; j++{
			pointArray = make([]float64, 2)
			pointArray[0]=p.rings[i].pos[j].x
			pointArray[1]=p.rings[i].pos[j].y
			ringArray=append(ringArray,pointArray)
		}
		ringArray=append(ringArray,ringArray[0])
		polygonArray = append(polygonArray, ringArray)
	}
	geomArray=append(geomArray, polygonArray)
	mj["coordinates"]= geomArray
	return mj
}

func (p Polygon) ExportGeoJSON() string {
	s,_ :=json.Marshal(p.ExportMap())
	return string(s)
}
