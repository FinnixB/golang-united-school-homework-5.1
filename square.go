package square

type Point struct {
	x, y int
}

func NewPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

type Square struct {
	start Point
	a     uint
}

func NewSquare(start Point, a uint) Square {
	return Square{start: start, a: a}
}

func (s Square) End() Point {
	return Point{x: s.start.x + int(s.a), y: s.start.y + int(s.a)}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
