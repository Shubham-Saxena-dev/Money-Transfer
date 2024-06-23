package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	var account *BankAccount
	BeforeEach(func() {
		account = new(BankAccount)
		account.ID = 1
		account.BalanceCents = 100
		account.BIC = "BIC"
		account.IBAN = "IBAN"
		account.OrganizationName = "ABC"

	})
	It("should create table and record in it", func() {
		err := db.AutoMigrate(BankAccount{})
		Ω(err).To(BeNil())
		err = db.Create(account).Error
		Ω(err).To(BeNil())
	})
})
