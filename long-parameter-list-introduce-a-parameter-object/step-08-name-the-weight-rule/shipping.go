package lesson

type Address struct {
	Postal, Country string
}

type Parcel struct {
	WeightKg                    float64
	LengthCm, WidthCm, HeightCm float64
}

type ShippingService struct {
	rates RateTable
}

func (s ShippingService) Quote(
	from, to Address, parcel Parcel, express bool,
) ShippingQuote {
	price := s.price(from, to, parcel, express)
	return ShippingQuote{from.Postal, to.Postal, price}
}

func (s ShippingService) price(
	from, to Address, parcel Parcel, express bool,
) float64 {
	cost := s.rates.PerKg(from.Country, to.Country) *
		s.billableKg(parcel)
	if express {
		cost *= 1.5
	}
	return cost
}

func (s ShippingService) billableKg(parcel Parcel) float64 {
	volumetricKg :=
		parcel.LengthCm * parcel.WidthCm * parcel.HeightCm / 5000
	return max(parcel.WeightKg, volumetricKg)
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
	parcel := Parcel{cart.WeightKg, cart.LengthCm,
		cart.WidthCm, cart.HeightCm}
	return c.shipping.Quote(from, to, parcel, cart.Express)
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

