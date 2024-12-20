package graphics_test

import (
	"math"
	"recal/cal"
	"recal/cal/graphics"
	"testing"
)

func TestGraph(t *testing.T) {
	x0 := 0.6
	var graph graphics.Graph
	var o graphics.Point
	for i := 0; i < 3; i++ {
		var p graphics.Point
		p[i] = 2.0
		graph.Arrows = append(graph.Arrows, graphics.Arrow{graphics.Line{o, p}})
	}
	f := cal.Function{
		Interval: [2]float64{0, 1.3},
		Formula:  cal.Cos,
	}
	g := cal.Function{
		Interval: [2]float64{0, 1.3},
		Formula:  cal.Sin,
	}
	h := cal.Function{
		Interval: [2]float64{0, 1.3},
		Formula:  func(x float64) float64 { return g.Formula(f.Formula(x)) },
	}
	fgph := f.Graph()
	for _, l := range fgph.Lines {
		graph.Lines = append(graph.Lines, graphics.Line{
			graphics.Point{l.From.X, l.From.Y, 0},
			graphics.Point{l.To.X, l.To.Y, 0},
		})
	}
	ggph := g.Graph()
	for _, l := range ggph.Lines {
		graph.Lines = append(graph.Lines, graphics.Line{
			graphics.Point{0, l.From.X, l.From.Y},
			graphics.Point{0, l.To.X, l.To.Y},
		})
	}
	hgph := h.Graph()
	for _, l := range hgph.Lines {
		graph.Lines = append(graph.Lines, graphics.Line{
			graphics.Point{l.From.X, 0, l.From.Y},
			graphics.Point{l.To.X, 0, l.To.Y},
		})
	}
	graph.Points = append(graph.Points,
		graphics.Point{x0, 0, 0},
		graphics.Point{x0, f.Formula(x0), 0},
		graphics.Point{0, f.Formula(x0), 0},
		graphics.Point{0, f.Formula(x0), g.Formula(f.Formula(x0))},
		graphics.Point{x0, 0, h.Formula(x0)})

	graph.Lines = append(graph.Lines,
		graphics.Line{
			graphics.Point{x0, 0, 0},
			graphics.Point{x0, f.Formula(x0), 0},
		},
		graphics.Line{
			graphics.Point{x0, f.Formula(x0), 0},
			graphics.Point{0, f.Formula(x0), 0},
		},
		graphics.Line{
			graphics.Point{0, f.Formula(x0), 0},
			graphics.Point{0, f.Formula(x0), h.Formula(x0)},
		},
		graphics.Line{
			graphics.Point{0, f.Formula(x0), h.Formula(x0)},
			graphics.Point{0, 0, h.Formula(x0)},
		},
		graphics.Line{
			graphics.Point{0, 0, h.Formula(x0)},
			graphics.Point{x0, 0, h.Formula(x0)},
		},
		graphics.Line{
			graphics.Point{x0, 0, h.Formula(x0)},
			graphics.Point{x0, 0, 0},
		},
	)
	skip, len := .05, .2
	graph.Arrows = append(graph.Arrows,
		graphics.Arrow{
			graphics.Line{
				graphics.Point{x0 + skip, -skip, 0},
				graphics.Point{x0 + len + skip, -skip, 0},
			},
		},
		graphics.Arrow{
			graphics.Line{
				graphics.Point{skip, f.Formula(x0) - skip, 0},
				graphics.Point{skip, f.Formula(x0) - len - skip, 0},
			},
		},
		graphics.Arrow{
			graphics.Line{
				graphics.Point{skip, 0, h.Formula(x0) - skip},
				graphics.Point{skip, 0, h.Formula(x0) - len - skip},
			},
		})
	graph.Rotate(5*math.Pi/4, -math.Pi/4)
	xy := graphics.XY{
		cal.NewXY(),
	}
	xy.Xlim = [2]float64{-2, 2}
	xy.Ylim = [2]float64{-2, 2}
	xy.Draw(graph)
	xy.Save()
}

func TestGraph_stokes(*testing.T) {
	var graph graphics.Graph
	g := cal.UnitCircle.Graph()
	for _, l := range g.Lines {
		graph.Lines = append(graph.Lines, graphics.Line{
			graphics.Point{l.From.X, l.From.Y, 0},
			graphics.Point{l.To.X, l.To.Y, 0}})
	}
	graph.Arrows = append(graph.Arrows,
		graphics.Arrow{
			graphics.Line{
				graphics.Point{0, 0, 0},
				graphics.Point{0, 0, 1},
			},
		})
	graph.Rotate(0, -1.3)
	graph.Rotate(1, 0)
	xy := graphics.XY{
		cal.NewXY(),
	}
	xy.Xlim = [2]float64{-1.2, 1.2}
	xy.Ylim = [2]float64{-1.2, 1.2}
	xy.Draw(graph)
	xy.Save()
}
