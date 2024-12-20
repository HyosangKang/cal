package cal

import (
	"image"
)

type XY struct {
	W, H       int             // width, height of the image
	Xlim, Ylim [2]float64      // x, y limits [min, max]
	Image      *image.Paletted // image pointer
}
