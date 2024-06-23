package customerrors

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errors", func() {
	Describe("RepositoryErrors", func() {
		It("should create a repositoryErrors and format the error message correctly", func() {
			err := errors.New("error")
			repoErr := ErrAcc(err)
			Ω(repoErr.Error()).To(Equal("error: account does not exit"))
			repoErr = ErrSaveRepo(err)
			Ω(repoErr.Error()).To(Equal("error: failed to save in repo"))
			repoErr = ErrUpdateRepo(err)
			Ω(repoErr.Error()).To(Equal("error: failed to update from repo"))
			repoErr = ErrBalance()
			Ω(repoErr.Error()).To(ContainSubstring("insufficient balance"))
		})
	})
})
