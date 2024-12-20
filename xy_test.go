package cal_test

import (
	"recal/cal"
	"testing"
)

func Linspace(x0, x1 float64, n int) []float64 {
	var arr []float64
	for i := 0; i <= n; i++ {
		x := x0 + float64(i)*(x1-x0)/float64(n)
		arr = append(arr, x)
	}
	return arr
}

func TestXY(*testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.5, 1.5}
	xy.Ylim = [2]float64{0.0, 3.0}
	xarr := Linspace(-1.0, 1.0, 10)
	for _, x := range xarr {
		xy.Point(x, cal.Exp(x))
	}
	xy.Save()
}
