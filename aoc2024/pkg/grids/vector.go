package grids

type Point struct {
	X, Y int
}

func P(x, y int) Point {
	return Point{X: x, Y: y}
}

type Vec2 struct {
	X, Y   int
	Dx, Dy int
}

func V2(x, y, dx, dy int) Vec2 {
	return Vec2{
		X: x, Y: y,
		Dx: dx, Dy: dy,
	}
}

func Step(v Vec2) Vec2 {
	return Vec2{v.X + v.Dx, v.Y + v.Dy, v.Dx, v.Dy}
}

func RotateRight(v Vec2) Vec2 {
	return Vec2{v.X, v.Y, v.Dy, -v.Dx}
}

func RotateLeft(v Vec2) Vec2 {
	return Vec2{v.X, v.Y, -v.Dy, v.Dx}
}
