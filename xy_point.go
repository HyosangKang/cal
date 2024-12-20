package cal

func (xy XY) Point(x, y float64) {
	i0, j0 := xy.Pixel(x, y)
	for i := -2; i <= 2; i++ {
		for j := -1; j <= 1; j++ {
			xy.Image.Set(i0+i, j0+j, Red)
			xy.Image.Set(i0+j, j0+i, Red)
		}
	}
}
