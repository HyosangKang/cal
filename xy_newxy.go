package cal

import (
	"image"
	"image/color"
)

const Width, Height = 400, 400

var (
	White = color.White
	Black = color.Black
	Red   = color.RGBA{255, 0, 0, 255}
)

func NewXY() XY {
	return XY{
		W: Width,
		H: Height,
		Image: image.NewPaletted(
			image.Rect(0, 0, Width, Height),
			color.Palette{
				White,
				Black,
				Red,
			}),
	}
}
