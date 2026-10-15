package lesson

import "fmt"

const (
	FreeShippingFrom = 50.0
	Fee              = 4.99
)

func ShipsFree(cart Cart) bool {
	return cart.Subtotal >= FreeShippingFrom
}

type Checkout struct{}

func (Checkout) ShippingFee(cart Cart) float64 {
	if ShipsFree(cart) {
		return 0.0
	}
	return Fee
}

type CartBanner struct{}

func (CartBanner) Message(cart Cart) string {
	if !ShipsFree(cart) {
		gap := FreeShippingFrom - cart.Subtotal
		return fmt.Sprintf("Add $%.2f to ship free", gap)
	}
	return "Your order ships free"
}

type OrderEmail struct{}

func (OrderEmail) ShippingLine(cart Cart) string {
	free := ShipsFree(cart)
	if free {
		return "Shipping: free"
	}
	return fmt.Sprintf("Shipping: $%.2f", Fee)
}
type Cart struct {
	Subtotal float64
	Member   bool
}
