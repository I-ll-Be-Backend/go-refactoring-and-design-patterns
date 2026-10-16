package lesson

import (
	"math"
	"time"
)

const (
	freeCancellationDays = 14
	lastMinuteDays       = 2
	nonRefundable        = 'N'
)

func CancellationFee(b Booking, today time.Time) float64 {
	daysBeforeCheckIn := int(b.CheckIn.Sub(today).Hours() / 24)
	if isFreeCancellation(daysBeforeCheckIn) {
		return 0
	}
	if paysInFull(b, daysBeforeCheckIn) {
		return b.Total
	}
	// otherwise half the total
	f := b.Total * 0.5
	// partner contracts require a one-night minimum
	return math.Max(f, b.Total/float64(b.Nights))
}

func isFreeCancellation(days int) bool {
	return days >= freeCancellationDays
}

func paysInFull(b Booking, days int) bool {
	return b.Rate == nonRefundable || days < lastMinuteDays
}
type Booking struct {
	CheckIn time.Time
	Total   float64
	Nights  int
	Rate    rune
	Loyal   bool
}
