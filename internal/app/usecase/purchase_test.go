package usecase_test

import (
	"blue-discount/internal/app"
	"blue-discount/internal/app/usecase"
	"blue-discount/internal/domain/model"
	repo_mocks "blue-discount/internal/domain/model/mocks"
	"blue-discount/internal/infra/repository"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	userID             = "ec1400b2-d1c3-471a-9fa3-c533a1ef1a3a"
	productID          = "15b69bd4-0357-45ed-96ca-62c308ed77e8"
	blackFridayPercent = int32(10)
	birthdayPercent    = int32(5)
	maxDiscount        = int32(10)
)

func newUser(bornAt *time.Time) model.User {
	ID, err := uuid.FromString(userID)
	Ω(err).Should(BeNil())
	return model.NewUser(ID, bornAt)
}

func newProduct(price int64) model.Product {
	ID, err := uuid.FromString(productID)
	Ω(err).Should(BeNil())
	return model.NewProduct(ID, price)
}

var _ = Describe("purchase usecase", func() {
	var (
		mockCtrl         *gomock.Controller
		userRepoMock     *repo_mocks.MockUserRepo
		productRepoMock  *repo_mocks.MockProductRepo
		campaignRepoMock *repo_mocks.MockCampaignRepo
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		userRepoMock = repo_mocks.NewMockUserRepo(mockCtrl)
		productRepoMock = repo_mocks.NewMockProductRepo(mockCtrl)
		campaignRepoMock = repo_mocks.NewMockCampaignRepo(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Discount()", func() {
		Context("when product price is 4500", func() {
			Context("and is black friday and birthday`s user", func() {
				It("should get 450 of discount", func() {
					now := time.Now()
					user := newUser(&now)

					userRepoMock.
						EXPECT().
						GetByID(userID).
						Return(user, nil)

					product := newProduct(4500)

					productRepoMock.
						EXPECT().
						GetByID(productID).
						Return(product, nil)

					campaigns := []model.Campaign{
						model.NewCampaign(model.BlackFridayCampaignName,
							blackFridayPercent, true, &now),

						model.NewCampaign(model.BirthdayCampaignName,
							birthdayPercent, true, nil),
					}

					campaignRepoMock.
						EXPECT().
						FindByActive(true).
						Return(campaigns, nil)

					cfg := app.DiscountConfig{MaxApplied: maxDiscount}

					usecase := usecase.NewPurchaseUsecaseImpl(userRepoMock,
						productRepoMock, campaignRepoMock, cfg)

					purchase, err := usecase.Discount(userID, productID)

					Ω(err).Should(BeNil())
					Ω(purchase.Discount.Value).Should(Equal(int64(450)))
				})
			})
		})

		Context("when product price is 3700", func() {
			Context("and is not black friday and birthday`s user", func() {
				It("should get 185 of discount", func() {
					now := time.Now()
					tomorrow := now.Add(24 * time.Hour)
					user := newUser(&now)

					userRepoMock.
						EXPECT().
						GetByID(userID).
						Return(user, nil)

					product := newProduct(3700)

					productRepoMock.
						EXPECT().
						GetByID(productID).
						Return(product, nil)

					campaigns := []model.Campaign{
						model.NewCampaign(model.BlackFridayCampaignName,
							blackFridayPercent, true, &tomorrow),

						model.NewCampaign(model.BirthdayCampaignName,
							birthdayPercent, true, nil),
					}

					campaignRepoMock.
						EXPECT().
						FindByActive(true).
						Return(campaigns, nil)

					cfg := app.DiscountConfig{MaxApplied: maxDiscount}

					usecase := usecase.NewPurchaseUsecaseImpl(userRepoMock,
						productRepoMock, campaignRepoMock, cfg)

					purchase, err := usecase.Discount(userID, productID)

					Ω(err).Should(BeNil())
					Ω(purchase.Discount.Value).Should(Equal(int64(185)))
				})
			})
		})

		It("should not found user", func() {
			userRepoMock.
				EXPECT().
				GetByID(userID).
				Return(model.User{}, repository.ErrRowNotFound)

			cfg := app.DiscountConfig{}

			usecase := usecase.NewPurchaseUsecaseImpl(userRepoMock,
				productRepoMock, campaignRepoMock, cfg)

			_, err := usecase.Discount(userID, productID)

			Ω(errors.Is(err, repository.ErrRowNotFound)).Should(Equal(true))
			Ω(err.Error()).Should(Equal(fmt.Sprintf("find user %s: record not found", userID)))
		})

		It("should get not found product error", func() {
			userRepoMock.
				EXPECT().
				GetByID(userID).
				Return(newUser(nil), nil)

			productRepoMock.
				EXPECT().
				GetByID(productID).
				Return(model.Product{}, repository.ErrRowNotFound)

			cfg := app.DiscountConfig{}

			usecase := usecase.NewPurchaseUsecaseImpl(userRepoMock,
				productRepoMock, campaignRepoMock, cfg)

			_, err := usecase.Discount(userID, productID)

			Ω(errors.Is(err, repository.ErrRowNotFound)).Should(Equal(true))
			Ω(err.Error()).Should(Equal(fmt.Sprintf("find product %s: record not found", productID)))
		})

		It("should get error when finding campaigns", func() {
			userRepoMock.
				EXPECT().
				GetByID(userID).
				Return(newUser(nil), nil)

			productRepoMock.
				EXPECT().
				GetByID(productID).
				Return(newProduct(100), nil)

			campaignRepoMock.
				EXPECT().
				FindByActive(true).
				Return(nil, errors.New("some error"))

			cfg := app.DiscountConfig{}

			usecase := usecase.NewPurchaseUsecaseImpl(userRepoMock,
				productRepoMock, campaignRepoMock, cfg)

			_, err := usecase.Discount(userID, productID)

			Ω(err).ShouldNot(BeNil())
			Ω(err.Error()).Should(Equal("find campaigns: some error"))
		})
	})
})
