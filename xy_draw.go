package cal

func (xy XY) Draw(g Graph) {
	for _, p := range g.Points {
		xy.Point(p.X, p.Y)
	}
	for _, l := range g.Lines {
		xy.Line(l.From, l.To)
	}
}
