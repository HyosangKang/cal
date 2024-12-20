package cal_test

import (
	"image"
	"image/color"
	"recal/cal"
	"testing"
)

func TestXY_Point(*testing.T) {
	xy := cal.XY{
		W:    7,
		H:    7,
		Xlim: [2]float64{-1, 1},
		Ylim: [2]float64{-1, 1},
		Image: image.NewPaletted(
			image.Rect(0, 0, 7, 7),
			color.Palette{cal.White, cal.Red},
		),
	}
	xy.Point(0, 0)
	xy.Save()
}
