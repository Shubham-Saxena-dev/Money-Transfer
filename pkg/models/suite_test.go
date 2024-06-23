package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errors Suite")
}

var db *gorm.DB
var _ = Describe("Service layer", func() {

	BeforeSuite(func() {
		db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	})

	AfterSuite(func() {
		err := db.Migrator().DropTable(BankAccount{}, Transfer{})
		Ω(err).To(BeNil())
	})
})
