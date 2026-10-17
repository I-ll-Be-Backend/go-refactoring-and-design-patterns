package lesson

func Charge(t Ticket) float64 {
	if t.Minutes <= 15 {
		return 0
	}
	return min(t.Hours()*3.0, 20.0)
}
type Ticket struct {
	Minutes  int
	Electric bool
}

// Hours rounds the stay up to whole hours.
func (t Ticket) Hours() float64 { return float64((t.Minutes + 59) / 60) }
