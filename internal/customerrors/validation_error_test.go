package customerrors

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errors", func() {
	Describe("ValidationErrors", func() {
		It("should create a validationErrors and format the error message correctly", func() {
			repoErr := ErrOrgBIC("error")
			Ω(repoErr.Error()).To(Equal("error: invalid org BIC"))
			repoErr = ErrOrgIBAN("error")
			Ω(repoErr.Error()).To(Equal("error: invalid org IBAN"))
			repoErr = ErrTransfer("error")
			Ω(repoErr.Error()).To(Equal("error: invalid transfer"))
			repoErr = ErrValue("someValue")
			Ω(repoErr.Error()).To(ContainSubstring("invalid value"))
		})
	})
})
