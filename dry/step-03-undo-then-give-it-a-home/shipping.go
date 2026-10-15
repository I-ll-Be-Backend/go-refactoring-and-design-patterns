package lesson

import "fmt"

type Checkout struct{}

func (Checkout) ShippingFee(cart Cart) float64 {
	if cart.Subtotal >= 50.0 {
		return 0.0
	}
	return 4.99
}

type CartBanner struct{}

func (CartBanner) Message(cart Cart) string {
	if cart.Subtotal < 50.0 {
		gap := 50.0 - cart.Subtotal
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
