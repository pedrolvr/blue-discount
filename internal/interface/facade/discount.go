package facade

import (
	"context"

	"blue-discount/internal/app/usecase"
	"blue-discount/internal/interface/dto"

	"github.com/go-kit/kit/endpoint"
)

func MakeCalculateDiscountEndpoint(usc usecase.PurchaseUsecase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CalculateDiscountDTO)
		purchase, err := usc.Discount(req.UserID, req.ProductID)

		return dto.PurchaseDTO{
			UserID:    purchase.User.ID.String(),
			ProductID: purchase.Product.ID.String(),
			Discount: dto.DiscountDTO{
				Percent: purchase.Discount.Percent,
				Value:   purchase.Discount.Value,
			},
			Err: err,
		}, nil
	}
}
