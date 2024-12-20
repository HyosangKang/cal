package cal

func (fn Function) GraphContinuous(eps float64) Graph {
	graph := fn.Graph()
	var lines []Line
	var points []Point
	for _, l := range graph.Lines {
		if Abs(l.From.Y-l.To.Y) > eps {
			points = append(points, l.From, l.To)
		} else {
			lines = append(lines, l)
		}
	}
	return Graph{
		Points: points,
		Lines:  lines,
	}
}
