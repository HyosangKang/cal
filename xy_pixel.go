package cal

func (xy XY) Pixel(x, y float64) (int, int) {
	a, b := xy.Xlim[0], xy.Xlim[1]
	c, d := xy.Ylim[0], xy.Ylim[1]
	dw := (b - a) / float64(xy.W)
	dh := (d - c) / float64(xy.H)
	i := int((x - a) / dw)
	j := int((d - y) / dh)
	return i, j
}
