package cal

func (c Curve) Graph() Graph {
	var graph Graph
	tarr := Linspace(c.Interval[0], c.Interval[1], Nsub)
	for i := 0; i < Nsub; i++ {
		t0, t1 := tarr[i], tarr[i+1]
		graph.Lines = append(graph.Lines, Line{
			From: Point{X: c.X(t0), Y: c.Y(t0)},
			To:   Point{X: c.X(t1), Y: c.Y(t1)},
		})
	}
	return graph
}
