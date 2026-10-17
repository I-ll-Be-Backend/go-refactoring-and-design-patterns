package lesson

func fee(t Ticket) float64 {
	f := 0.0
	f = t.Hours() * 3.0
	f = min(f, 20.0)
	if t.Minutes <= 15 {
		f = 0
	}
	return f
}

func Charge(t Ticket) float64 {
	return fee(t)
}
type Ticket struct {
	Minutes  int
	Electric bool
}

// Hours rounds the stay up to whole hours.
func (t Ticket) Hours() float64 { return float64((t.Minutes + 59) / 60) }
