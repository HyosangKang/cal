package cal

func (xy XY) Line(p, q Point) {
	i0, j0 := xy.Pixel(p.X, p.Y)
	i1, j1 := xy.Pixel(q.X, q.Y)
	n := Max(Abs(i0-i1), Abs(j0-j1))
	is := Linspace(i0, i1, n)
	js := Linspace(j0, j1, n)
	for k := 0; k <= n; k++ {
		i2, j2 := is[k], js[k]
		for i := -1; i <= 1; i++ {
			xy.Image.Set(i2+i, j2, Black)
			xy.Image.Set(i2, j2+i, Black)
		}
	}
}

func Max[T float64 | int](n, m T) T {
	if n < m {
		return m
	}
	return n
}

func Abs[T float64 | int](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Linspace[T float64 | int](a, b T, n int) []T {
	af, bf := float64(a), float64(b)
	h := (bf - af) / float64(n)
	var arr []T
	for i := 0; i <= n; i++ {
		arr = append(arr, T(af+float64(i)*h))
	}
	return arr
}
