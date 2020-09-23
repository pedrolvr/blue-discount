package usecase

import (
	"blue-discount/internal/app"
	"blue-discount/internal/domain/model"
	"fmt"
)

type PurchaseUsecase interface {
	Discount(userID, prodcutID string) (model.Purchase, error)
}

type PurchaseUsecaseImpl struct {
	userRepo       model.UserRepo
	productRepo    model.ProductRepo
	campaignRepo   model.CampaignRepo
	discountConfig app.DiscountConfig
}

func NewPurchaseUsecaseImpl(u model.UserRepo, p model.ProductRepo, c model.CampaignRepo, cfg app.DiscountConfig) PurchaseUsecase {
	return &PurchaseUsecaseImpl{
		userRepo:       u,
		productRepo:    p,
		campaignRepo:   c,
		discountConfig: cfg,
	}
}

func (u *PurchaseUsecaseImpl) Discount(userID, productID string) (model.Purchase, error) {
	var purchase model.Purchase

	user, err := u.userRepo.GetByID(userID)

	if err != nil {
		errMsg := "find user %s: %w"
		return purchase, fmt.Errorf(errMsg, userID, err)
	}

	product, err := u.productRepo.GetByID(productID)

	if err != nil {
		errMsg := "find product %s: %w"
		return purchase, fmt.Errorf(errMsg, productID, err)
	}

	campaigns, err := u.campaignRepo.FindByActive(true)

	if err != nil {
		errMsg := "find campaigns: %w"
		return purchase, fmt.Errorf(errMsg, err)
	}

	purchase = model.NewPurchase(user, product, model.NewDiscount(campaigns))
	purchase.CalculateDiscount(u.discountConfig.MaxApplied)

	return purchase, nil
}
