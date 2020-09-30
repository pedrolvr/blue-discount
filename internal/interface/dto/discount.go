package dto

type CalculateDiscountDTO struct {
	UserID    string
	ProductID string
}

type DiscountDTO struct {
	Percent int32
	Value   int64
}

type PurchaseDTO struct {
	UserID    string
	ProductID string
	Discount  DiscountDTO
}
