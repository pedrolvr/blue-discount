package model

type Purchase struct {
	User     User     `json:"user"`
	Product  Product  `json:"product"`
	Discount Discount `json:"discount"`
}

func NewPurchase(user User, product Product, discount Discount) Purchase {
	return Purchase{
		User:     user,
		Product:  product,
		Discount: discount,
	}
}
