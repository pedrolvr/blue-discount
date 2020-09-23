package model_test

import (
	"blue-discount/internal/domain/model"
	"time"

	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("user test", func() {
	Describe("NewUser()", func() {
		It("should get user", func() {
			now := time.Now()
			uID := uuid.Must(uuid.NewV4())
			user := model.NewUser(uID, &now)

			Ω(user.ID).Should(Equal(uID))
			Ω(user.BornAt).Should(Equal(&now))
		})
	})
})
