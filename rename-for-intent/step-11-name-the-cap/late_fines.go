package lesson

import "fmt"

type Loan struct {
	DaysOut int
	Student bool
}

const (
	loanDays     = 14
	studentRate  = 0.10
	standardRate = 0.25
	maxFine      = 10.0
)

func calc(loan Loan) float64 {
	daysLate := loan.DaysOut - loanDays
	if daysLate <= 0 {
		return 0
	}
	dailyRate := standardRate
	if loan.Student {
		dailyRate = studentRate
	}
	return min(float64(daysLate)*dailyRate, maxFine)
}

func receipt(loan Loan) string {
	return fmt.Sprintf("Late fine: $%.2f", calc(loan))
}

