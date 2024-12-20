package cal

const Eps = 1e-10

func IsZero(x float64) bool {
	return -Eps < x && x < Eps
}
