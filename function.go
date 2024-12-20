package cal

type Function struct {
	Interval [2]float64
	Formula  func(float64) float64
}
