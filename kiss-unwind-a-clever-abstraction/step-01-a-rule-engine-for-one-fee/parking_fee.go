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

func fee(t Ticket, plan []string) float64 {
	f := 0.0
	for _, name := range plan {
		f = rules[name](t, f)
	}
	return f
}

func Charge(t Ticket) float64 {
	return fee(t, []string{"base", "cap", "grace"})
}
type Ticket struct {
	Minutes  int
	Electric bool
}

// Hours rounds the stay up to whole hours.
func (t Ticket) Hours() float64 { return float64((t.Minutes + 59) / 60) }
