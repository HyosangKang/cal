package cal_test

import (
	"recal/cal"
	"testing"
)

func TestGraph_linear(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{0, 3}
	exp := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  cal.Exp,
	}
	xy.Draw(exp.Graph())
	f := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  func(x float64) float64 { return x + 1 },
	}
	xy.Draw(f.Graph())
	xy.Save()
}

func TestGraph_quadratic(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{0, 3}
	exp := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  cal.Exp,
	}
	xy.Draw(exp.Graph())
	f := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			return x*x/2 + x + 1
		},
	}
	xy.Draw(f.Graph())
	xy.Save()
}

func TestGraph_cubic(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{0, 3}
	exp := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  cal.Exp,
	}
	xy.Draw(exp.Graph())
	f := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			return x*x*x/6 + x*x/2 + x + 1
		},
	}
	xy.Draw(f.Graph())
	xy.Save()
}

func TestGraph_log(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{.1, 3.9}
	xy.Ylim = [2]float64{-1, 3}
	log := cal.Function{
		Interval: [2]float64{.1, 3.9},
		Formula:  cal.Log,
	}
	g := cal.Function{
		Interval: [2]float64{.1, 3.9},
		Formula: func(x float64) float64 {
			var f, p float64 = 0, (x - 1)
			for i := 0; i < 5; i++ {
				f += p
				s := float64(i)
				p *= -(s + 1) / (s + 2) * (x - 1)
			}
			return f
		},
	}
	xy.Draw(log.Graph())
	xy.Draw(g.Graph())
	xy.Save()
}
