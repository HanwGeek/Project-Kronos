package geom

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

	l2:=NewLayer("test2",KrPolygon)
	po:=Polygon{}
	outr:=LineString{[]Coord{Coord{1,2},Coord{4,7},Coord{-3,-3}}}
	innr:=LineString{[]Coord{Coord{0.5,0.5},Coord{0.3,0.3},Coord{0.3,0.5}}}
	po.AddRing(outr)
	po.AddRing(innr)
	l2.AddFeature(*NewFeatureByGeom(po))
	l2.AddFeature(*NewFeatureByGeom(p))
	fmt.Println(l2.FeatureCount())
	fmt.Println(l2.GetFeature(0))
	npo:=l2.GetFeature(0).geom.(Polygon)
	fmt.Println(npo.NInnerRings())
	npo.DeleteRing(1)
	fmt.Println(npo)
}
*/
