package lesson

type Rule func(t Ticket, fee float64) float64

var rules = map[string]Rule{
	"base": func(t Ticket, f float64) float64 {
		return t.Hours() * 3.0
	},
	"cap": func(t Ticket, f float64) float64 {
		return min(f, 20.0)
	},
	"grace": func(t Ticket, f float64) float64 {
		if t.Minutes <= 15 {
			return 0
		}
		return f
	},
}

func fee(t Ticket) float64 {
	f := 0.0
	f = rules["base"](t, f)
	f = rules["cap"](t, f)
	f = rules["grace"](t, f)
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
