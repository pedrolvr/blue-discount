package model

type Purchase struct {
	User         User               `json:"user"`
	Product      Product            `json:"product"`
	Discount     Discount           `json:"discount"`
	DiscountCalc DiscountCalculator `json:"-"`
}

func NewPurchase(user User, product Product, discountCalc DiscountCalculator) Purchase {
	return Purchase{
		User:         user,
		Product:      product,
		DiscountCalc: discountCalc,
	}
}

func (m *Purchase) CalculateDiscount(maxPercent int32) {
	m.Discount = m.DiscountCalc.Calculate(maxPercent, m.User, m.Product)
}
