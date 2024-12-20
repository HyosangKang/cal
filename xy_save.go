package cal

import (
	"image/png"
	"os"
)

func (xy XY) Save() {
	fp, _ := os.Create("image.png")
	defer fp.Close()
	png.Encode(fp, xy.Image)
}
