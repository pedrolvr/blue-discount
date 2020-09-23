package repository_test

import (
	"database/sql"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

func connectWithDB(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{Conn: db})
	dbConn, err := gorm.Open(dialector, &gorm.Config{})
	Ω(err).ShouldNot(HaveOccurred())
	return dbConn
}
