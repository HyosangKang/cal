package graphics

import (
	"math"
	"recal/cal"
)

type XY struct {
	cal.XY
}

const Tip = 5

func (xy XY) Arrow(a cal.Line) {
	xy.Line(a.From, a.To)
	i0, j0 := xy.Pixel(a.To.X, a.To.Y)
	i1, j1 := xy.Pixel(a.From.X, a.From.Y)
	di, dj := float64(i0-i1), float64(j0-j1)
	d := math.Sqrt(di*di + dj*dj)
	p := cal.Point{
		X: (a.From.X - a.To.X) * Tip / d,
		Y: (a.From.Y - a.To.Y) * Tip / d}
	for _, t := range []float64{-math.Pi / 4, math.Pi / 4} {
		c, s := cal.Cos(t), cal.Sin(t)
		xy.Line(a.To, cal.Point{
			X: a.To.X + p.X*c - p.Y*s,
			Y: a.To.Y + p.X*s + p.Y*c})
	}
}

func (xy XY) Draw(g Graph) {
	for _, p := range g.Points {
		q := p.Project()
		xy.Point(q.X, q.Y)
	}
	for _, l := range g.Lines {
		m := l.Project()
		xy.Line(m.From, m.To)
	}
	for _, a := range g.Arrows {
		b := a.Project()
		xy.Arrow(b)
	}
}
