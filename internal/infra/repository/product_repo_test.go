package repository_test

import (
	"blue-discount/internal/domain/model"
	"blue-discount/internal/infra/repository"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("product repo", func() {
	var repo model.ProductRepo
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		repo = repository.NewProductRepo(connectWithDB(db))
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("GetByID()", func() {
		ID, err := uuid.FromString("f57d2888-4536-413d-8c3f-760f75a10232")
		Expect(err).ShouldNot(HaveOccurred())

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product"`)).
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(ID.String()),
			)

		row, err := repo.GetByID(ID.String())

		Expect(err).ShouldNot(HaveOccurred())
		Expect(row.ID).Should(Equal(ID))
	})
})
