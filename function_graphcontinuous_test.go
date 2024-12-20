package cal_test

import (
	"recal/cal"
	"testing"
)

func TestFunction_Graph_dirac(*testing.T) {
	f := cal.Function{
		Interval: [2]float64{-1., 1.},
		Formula: func(x float64) float64 {
			if x < 0 {
				return -1
			} else if x > 0 {
				return 1
			} else {
				return 0
			}
		}}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(f.Graph())
	xy.Save()
}

func TestFunction_Graph_dirac_two(*testing.T) {
	f0 := cal.Function{
		Interval: [2]float64{-1., 0.},
		Formula: func(x float64) float64 {
			return -1
		}}
	f1 := cal.Function{
		Interval: [2]float64{0., 1.},
		Formula: func(x float64) float64 {
			return 1
		}}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(f0.Graph())
	xy.Draw(f1.Graph())
	xy.Save()
}

func TestFunction_Graph_delta(*testing.T) {
	f := cal.Function{
		Interval: [2]float64{-1., 1.},
		Formula: func(x float64) float64 {
			return 0
		}}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-.1, 1.1}
	xy.Draw(f.Graph())
	xy.Point(0, 1)
	xy.Save()
}
