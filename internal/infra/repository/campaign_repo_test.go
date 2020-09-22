package repository_test

import (
	"blue-discount/internal/domain/model"
	"blue-discount/internal/infra/repository"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("campaign repo", func() {
	var repo model.CampaignRepo
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		repo = repository.NewCampaignRepo(connectWithDB(db))
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("FindByActive()", func() {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "campaign"`)).
			WillReturnRows(
				sqlmock.NewRows([]string{"name"}).AddRow(model.BirthdayCampaignName),
			)

		rows, err := repo.FindByActive(true)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(rows[0].Name).Should(Equal(model.BirthdayCampaignName))
	})
})
