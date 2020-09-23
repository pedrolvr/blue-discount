package model_test

import (
	"blue-discount/internal/domain/model"

	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("product test", func() {
	Describe("NewProduct()", func() {
		It("should get product", func() {
			price := int64(3700)
			uID := uuid.Must(uuid.NewV4())
			product := model.NewProduct(uID, price)

			Ω(product.ID).Should(Equal(uID))
			Ω(product.Price).Should(Equal(price))
		})
	})
})
