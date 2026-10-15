package lesson

import "fmt"

type Loan struct {
	DaysOut int
	Student bool
}

const (
	loanDays = 14
)

func calc(loan Loan) float64 {
	daysLate := loan.DaysOut - loanDays
	if daysLate <= 0 {
		return 0
	}
	dailyRate := 0.25
	if loan.Student {
		dailyRate = 0.10
	}
	return min(float64(daysLate)*dailyRate, 10.0)
}

func receipt(loan Loan) string {
	return fmt.Sprintf("Late fine: $%.2f", calc(loan))
}

