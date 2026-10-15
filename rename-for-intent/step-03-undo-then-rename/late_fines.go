package lesson

import "fmt"

type Loan struct {
	D int
	F bool
}

func calc(l Loan) float64 {
	x := l.D - 14
	if x <= 0 {
		return 0
	}
	r := 0.25
	if l.F {
		r = 0.10
	}
	return min(float64(x)*r, 10.0)
}

func receipt(l Loan) string {
	return fmt.Sprintf("Late fine: $%.2f", calc(l))
}

