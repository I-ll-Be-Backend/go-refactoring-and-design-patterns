package lesson

type ShippingService struct {
	rates RateTable
}

func (s ShippingService) Quote(
	fromPostal, fromCountry, toPostal, toCountry string,
	weightKg, lengthCm, widthCm, heightCm float64,
	express, fragile bool,
) ShippingQuote {
	price := s.price(fromCountry, toCountry,
		weightKg, lengthCm, widthCm, heightCm, express, fragile)
	return ShippingQuote{fromPostal, toPostal, price}
}

func (s ShippingService) price(
	fromCountry, toCountry string,
	weightKg, lengthCm, widthCm, heightCm float64,
	express, fragile bool,
) float64 {
	volumetricKg := lengthCm * widthCm * heightCm / 5000
	cost := s.rates.PerKg(fromCountry, toCountry) *
		max(weightKg, volumetricKg)
	if fragile {
		cost *= 1.2
	}
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
	return c.shipping.Quote(warehousePostal, warehouseCountry,
		customer.PostalCode, customer.Country,
		cart.WeightKg, cart.LengthCm, cart.WidthCm, cart.HeightCm,
		cart.Express, cart.Fragile)
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

