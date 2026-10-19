package lesson

type Address struct {
	Postal, Country string
}

type Parcel struct {
	WeightKg                    float64
	LengthCm, WidthCm, HeightCm float64
}

func (p Parcel) billableKg() float64 {
	volumetricKg :=
		p.LengthCm * p.WidthCm * p.HeightCm / 5000
	return max(p.WeightKg, volumetricKg)
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
		parcel.billableKg()
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

