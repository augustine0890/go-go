package main

func GetDisplaceFn(a, b, c float64) func(float64) float64 {
	fn := func(d float64) float64 {
		return (0.5)*(a*d*d) + (b * d) + c
	}
	return fn
}
