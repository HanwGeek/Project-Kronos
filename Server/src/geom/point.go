package geom

type Point struct {
	pos Coord
}

func (p Point) GetX() float32 {
	return p.pos.x
}

func (p Point) GetY() float32 {
	return p.pos.y
}

func (p Point) SetX(_x float32) {
	p.pos.x = _x
}

func (p Point) SetY(_y float32) {
	p.pos.y = _y
}

/*
func main(){
	var p,q Point
	p.SetX(10)
	p.SetY(10)
	q=Point{Coord{5,5}}
	fmt.Println(q.GetX())
}
*/
