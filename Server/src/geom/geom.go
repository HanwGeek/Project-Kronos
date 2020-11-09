package geom

import (
	//"fmt"
	"strconv"
	"strings"
)

const (
	// KrPoint is a point
	KrPoint = 1
	// KrLineString is a linestring
	KrLineString = 2
	// KrLinerRing is a linear ring with one end point
	KrLinerRing = 3
	// KrPolygon is a polygon
	KrPolygon = 4
	// KrMultiPoint is a collection of points
	KrMultiPoint = 5
	// KrMultiLineString is a collection of linestrings
	KrMultiLineString = 6
	// KrMultiPolygon is a collection of polygons
	KrMultiPolygon = 7
	// KrGeometryCollection is a collection of heterogeneus geometries
	KrGeometryCollection = 8
)

type Geometry interface {
	GeomType() int
	ExportWKT() string
}

func NewGeomByWKT(str string) Geometry {
	if strings.Contains(str, "POINT") {
		f_ := func(c rune) bool {
			if c == ' ' || c == '(' || c == ')' {
				return true
			} else {
				return false
			}
		}
		segs := strings.FieldsFunc(str, f_)
		x_, _ := strconv.ParseFloat(segs[1], 64)
		y_, _ := strconv.ParseFloat(segs[2], 64)
		return Point{Coord{x_, y_}}

	} else if strings.Contains(str, "LINESTRING") {
		f_ := func(c rune) bool {
			if c == ' ' || c == '(' || c == ')' || c == ',' {
				return true
			} else {
				return false
			}
		}
		segs := strings.FieldsFunc(str, f_)
		newline := LineString{}
		for p := 0; 2*p+2 < len(segs); p++ {
			x_, _ := strconv.ParseFloat(segs[2*p+1], 64)
			y_, _ := strconv.ParseFloat(segs[2*p+2], 64)
			newline.AddPoint(x_, y_)
		}
		return newline
	} else if strings.Contains(str, "POLYGON") {
		f_ := func(c rune) bool {
			if c == ' ' || c == ')' || c == ',' {
				return true
			} else {
				return false
			}
		}
		segs := strings.FieldsFunc(str, f_)
		strp := strings.Join(segs[1:], " ")
		segs = strings.Split(strp, "(")[2:]
		newpoly := Polygon{}
		outr := LineString{}
		strout := strings.Fields(segs[0])
		for i := 0; 2*i+1 < len(strout)-2; i++ {
			x_, _ := strconv.ParseFloat(strout[2*i], 64)
			y_, _ := strconv.ParseFloat(strout[2*i+1], 64)
			outr.AddPoint(x_, y_)
		}
		newpoly.AddRing(outr)
		for r := 1; r < len(segs); r++ {
			innr := LineString{}
			strin := strings.Fields(segs[r])
			for i := 0; 2*i+1 < len(strin)-2; i++ {
				x_, _ := strconv.ParseFloat(strin[2*i], 64)
				y_, _ := strconv.ParseFloat(strin[2*i+1], 64)
				innr.AddPoint(x_, y_)
			}
			newpoly.AddRing(innr)
		}
		return newpoly
	}
	return nil
}

/*
func main(){
	l:=NewLayer("test1",KrPoint)
	fmt.Println(l.IsEmpty())
	p:=Point{Coord{3,5}}
	fmt.Println(p.GeomType())
	fmt.Println(l.GetGeomType())
	fp:=NewFeatureByGeom(p)
	fp.SetAttribute("Area",500)
	fp.SetAttribute("Name","Hanwgeek's Villa")
	l.AddFeature(*fp)
	fmt.Println(l.feat)
	fmt.Println(l.FeatureCount())
	fv:=l.GetFeature(0)
	fmt.Println(fv)
	fmt.Println(fv.geom.ExportWKT())

	l2:=NewLayer("test2",KrPolygon)
	po:=Polygon{}
	outr:=LineString{[]Coord{Coord{1,2},Coord{4,7},Coord{-3,-3}}}
	innr:=LineString{[]Coord{Coord{0.5,0.5},Coord{0.3,0.3},Coord{0.3,0.5}}}
	fmt.Println(outr.ExportWKT())
	fmt.Println(innr.ExportWKT())
	po.AddRing(outr)
	po.AddRing(innr)
	l2.AddFeature(*NewFeatureByGeom(po))
	l2.AddFeature(*NewFeatureByGeom(p))
	fmt.Println(l2.FeatureCount())
	fmt.Println(l2.GetFeature(0))
	npo:=l2.GetFeature(0).geom.(Polygon)
	fmt.Println(npo.NInnerRings())
	fmt.Println(npo.ExportWKT())
	npo.DeleteRing(1)
	fmt.Println(npo)

	newgeom:=NewGeomByWKT("POINT (3.000000 5.000000)")
	fmt.Println(newgeom)
	newgeom=NewGeomByWKT("LINESTRING (0.500000 0.500000,0.300000 0.300000,0.300000 0.500000)")
	fmt.Println(newgeom)
	newgeom=NewGeomByWKT("POLYGON ((1.000000 2.000000,4.000000 7.000000,-3.000000 -3.000000,1.000000 2.000000),(0.500000 0.500000,0.300000 0.300000,0.300000 0.500000,0.500000 0.500000))")
	fmt.Println(newgeom)
}
*/
