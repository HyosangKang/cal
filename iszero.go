package cal

const Eps = 1e-6

func IsZero(x float64) bool {
	return -Eps < x && x < Eps
}
