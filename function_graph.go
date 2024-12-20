package cal

const Nsub = 100

func (fn Function) Graph() Graph {
	a, b := fn.Interval[0], fn.Interval[1]
	xarr := Linspace(a, b, Nsub)
	var graph Graph
	for i := 0; i < Nsub; i++ {
		x0, x1 := xarr[i], xarr[i+1]
		y0, y1 := fn.Formula(x0), fn.Formula(x1)
		graph.Lines = append(graph.Lines, Line{
			From: Point{X: x0, Y: y0},
			To:   Point{X: x1, Y: y1},
		})
	}
	return graph
}
