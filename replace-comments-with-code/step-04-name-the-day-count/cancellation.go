package lesson

import (
	"math"
	"time"
)

func CancellationFee(b Booking, today time.Time) float64 {
	daysBeforeCheckIn := int(b.CheckIn.Sub(today).Hours() / 24)
	// free cancellation 14 days or more before check-in
	if daysBeforeCheckIn >= 14 {
		return 0
	}
	// non-refundable or under 2 days out: pay in full
	if b.Rate == 'N' || daysBeforeCheckIn < 2 {
		return b.Total
	}
	// otherwise half the total
	f := b.Total * 0.5
	// partner contracts require a one-night minimum
	return math.Max(f, b.Total/float64(b.Nights))
}
type Booking struct {
	CheckIn time.Time
	Total   float64
	Nights  int
	Rate    rune
	Loyal   bool
}
