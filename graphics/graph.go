package graphics

import "recal/cal"

type Graph struct {
	Points []Point
	Lines  []Line
	Arrows []Arrow
}

func (g *Graph) FromGraph(cg cal.Graph) {
	for _, p := range cg.Points {
		g.Points = append(g.Points, Point{p.X, p.Y, 0})
	}
	for _, l := range cg.Lines {
		g.Lines = append(g.Lines, Line{
			Point{l.From.X, l.From.Y, 0},
			Point{l.To.X, l.To.Y, 0},
		})
	}

}

func (g *Graph) Rotate(theta, phi float64) {
	for i, p := range g.Points {
		g.Points[i] = p.Rotate(theta, phi)
	}
	for i, l := range g.Lines {
		g.Lines[i] = l.Rotate(theta, phi)
	}
	for i, a := range g.Arrows {
		g.Arrows[i] = Arrow{a.Rotate(theta, phi)}
	}
}

type Point [3]float64

func (p Point) Rotate(theta, phi float64) Point {
	p = Rot(theta, [2]int{0, 1}).Mul(p)
	return Rot(phi, [2]int{1, 2}).Mul(p)
}

func (p Point) Project() cal.Point {
	return cal.Point{
		X: p[0],
		Y: p[1],
	}
}

type Line [2]Point

func (l Line) Rotate(theta, phi float64) Line {
	return Line{l[0].Rotate(theta, phi), l[1].Rotate(theta, phi)}
}

func (l Line) Project() cal.Line {
	return cal.Line{
		From: l[0].Project(),
		To:   l[1].Project(),
	}
}

type Arrow struct {
	Line
}
