package lesson

import "fmt"

type Checkout struct{}

func (Checkout) ShippingFee(cart Cart) float64 {
	freeFrom := 50.0
	if cart.Member {
		freeFrom = 25.0
	}
	if cart.Subtotal >= freeFrom {
		return 0.0
	}
	return 4.99
}

type CartBanner struct{}

func (CartBanner) Message(cart Cart) string {
	limit := 50.0
	if cart.Member {
		limit = 25.0
	}
	if cart.Subtotal < limit {
		gap := limit - cart.Subtotal
		return fmt.Sprintf("Add $%.2f to ship free", gap)
	}
	return "Your order ships free"
}

type OrderEmail struct{}

func (OrderEmail) ShippingLine(cart Cart) string {
	free := cart.Subtotal >= 50.0
	if free {
		return "Shipping: free"
	}
	return "Shipping: $4.99"
}
type Cart struct {
	Subtotal float64
	Member   bool
}
