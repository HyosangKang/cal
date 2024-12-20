package graphics

import "recal/cal"

type Mat [3][3]float64

func (m Mat) Mul(p Point) Point {
	var q Point
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			q[i] += m[i][j] * p[j]
		}
	}
	return q
}

func Rot(t float64, idx [2]int) Mat {
	m := Id()
	m[idx[0]][idx[0]] = cal.Cos(t)
	m[idx[0]][idx[1]] = -cal.Sin(t)
	m[idx[1]][idx[0]] = cal.Sin(t)
	m[idx[1]][idx[1]] = cal.Cos(t)
	return m
}

func Id() Mat {
	var m Mat
	for i := 0; i < 3; i++ {
		m[i][i] = 1
	}
	return m
}
