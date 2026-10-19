package lesson

type Address struct {
	Postal, Country string
}

type ShippingService struct {
	rates RateTable
}

func (s ShippingService) Quote(
	from, to Address,
	weightKg, lengthCm, widthCm, heightCm float64,
	express bool,
) ShippingQuote {
	price := s.price(from, to,
		weightKg, lengthCm, widthCm, heightCm, express)
	return ShippingQuote{from.Postal, to.Postal, price}
}

func (s ShippingService) price(
	from, to Address,
	weightKg, lengthCm, widthCm, heightCm float64,
	express bool,
) float64 {
	volumetricKg := lengthCm * widthCm * heightCm / 5000
	cost := s.rates.PerKg(from.Country, to.Country) *
		max(weightKg, volumetricKg)
	if express {
		cost *= 1.5
	}
	return cost
}

const (
	warehousePostal  = "10115"
	warehouseCountry = "DE"
)

type Checkout struct {
	shipping ShippingService
}

func (c Checkout) ShippingFor(cart Cart, customer Customer) ShippingQuote {
	from := Address{warehousePostal, warehouseCountry}
	to := Address{customer.PostalCode, customer.Country}
	return c.shipping.Quote(from, to,
		cart.WeightKg, cart.LengthCm, cart.WidthCm, cart.HeightCm,
		cart.Express)
}
type RateTable interface {
	PerKg(fromCountry, toCountry string) float64
}

type ShippingQuote struct {
	FromPostal, ToPostal string
	Price                float64
}

type Cart struct {
	WeightKg, LengthCm, WidthCm, HeightCm float64
	Express, Fragile                      bool
}

type Customer struct {
	PostalCode, Country string
}

