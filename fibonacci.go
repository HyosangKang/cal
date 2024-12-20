package cal

func Fibonacci(n int) float64 {
	if n == 0 || n == 1 {
		return 1
	}
	f := [2]float64{1, 1}
	for i := 2; i <= n; i++ {
		f[0], f[1] = f[1], f[0]+f[1]
	}
	return f[1]
}
