package cal

// func Plot(itv [2]float64, fn func(float64) float64) {
// 	xs := Linspace(itv, Nsub)
// 	ys := make([]float64, len(xs))
// 	for i, x := range xs {
// 		ys[i] = fn(x)
// 	}
// 	ymin, ymax := MinMax(ys)
// 	xy := XY{
// 		W:    400,
// 		H:    400,
// 		Xlim: itv,
// 		Ylim: [2]float64{ymin, ymax}}
// 	img := image.NewPaletted(
// 		image.Rect(0, 0, xy.W, xy.H),
// 		color.Palette{color.White, color.Black})
// 	for i := 0; i < Nsub; i++ {
// 		i0, j0 := xy.Pixel(xs[i], ys[i])
// 		i1, j1 := xy.Pixel(xs[i+1], ys[i+1])
// 		DrawLine(img, i0, j0, i1, j1)
// 	}
// 	fp, _ := os.Create("image.png")
// 	defer fp.Close()
// 	png.Encode(fp, img)
// }

func MinMax(ys []float64) (float64, float64) {
	var ymin, ymax float64 = ys[0], ys[0]
	for _, y := range ys {
		if y < ymin {
			ymin = y
		}
		if y > ymax {
			ymax = y
		}
	}
	return ymin, ymax
}
