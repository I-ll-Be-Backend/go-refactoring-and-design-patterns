package lesson

import "fmt"

type Loan struct {
	DaysOut int
	Student bool
}

func calc(l Loan) float64 {
	daysLate := l.DaysOut - 14
	if daysLate <= 0 {
		return 0
	}
	r := 0.25
	if l.Student {
		r = 0.10
	}
	return min(float64(daysLate)*r, 10.0)
}

func receipt(l Loan) string {
	return fmt.Sprintf("Late fine: $%.2f", calc(l))
}

