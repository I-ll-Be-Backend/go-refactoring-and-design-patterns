package lesson

const (
	hourlyRate   = 3.0
	dailyCap     = 20.0
	graceMinutes = 15
)

func Charge(t Ticket) float64 {
	if t.Minutes <= graceMinutes {
		return 0
	}
	return min(t.Hours()*hourlyRate, dailyCap)
}
type Ticket struct {
	Minutes  int
	Electric bool
}

// Hours rounds the stay up to whole hours.
func (t Ticket) Hours() float64 { return float64((t.Minutes + 59) / 60) }
