package repository_test

import (
	"blue-discount/internal/domain/model"
	"blue-discount/internal/infra/repository"
	"database/sql"

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
		Ω(err).ShouldNot(HaveOccurred())

		repo = repository.NewCampaignRepo(connectWithDB(db))
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("FindByActive()", func() {
		mock.ExpectQuery(`SELECT (.+) FROM "campaign" (.+) ORDER BY order ASC`).
			WillReturnRows(
				sqlmock.NewRows([]string{"name"}).AddRow(model.BirthdayCampaignName),
			)

		rows, err := repo.FindByActive(true)

		Ω(err).ShouldNot(HaveOccurred())
		Ω(rows[0].Name).Should(Equal(model.BirthdayCampaignName))
	})
})
