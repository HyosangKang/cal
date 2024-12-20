package cal

type Point struct {
	X, Y float64
}

type Line struct {
	From, To Point
}

type Graph struct {
	Points []Point
	Lines  []Line
}
