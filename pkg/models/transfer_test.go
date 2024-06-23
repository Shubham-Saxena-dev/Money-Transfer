package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transfer", func() {
	var transfer *Transfer
	BeforeEach(func() {
		transfer = new(Transfer)
		transfer.ID = 1
		transfer.AmountCents = 100
		transfer.BankAccountID = 1
		transfer.CounterpartyBIC = "BIC"
		transfer.CounterpartyIBAN = "IBAN"
		transfer.CounterpartyName = "ABC"
		transfer.Description = "DEF"
	})

	It("should create table and record in it", func() {
		err := db.AutoMigrate(Transfer{})
		Ω(err).To(BeNil())
		err = db.Create(transfer).Error
		Ω(err).To(BeNil())
	})
})
