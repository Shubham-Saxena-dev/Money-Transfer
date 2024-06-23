package repositories

import (
	"gorm.io/gorm"
	"qonto/internal/database"
	"qonto/pkg/models"
)

type BankAccountRepository interface {
	GetBalnceByIBANAndBIC(iban, bic string) (models.BankAccount, error)
	Update(account models.BankAccount) error
}

type bankAccountRepository struct {
	*gorm.DB
}

func NewBankAccountRepository() BankAccountRepository {
	return &bankAccountRepository{database.GetInstance()}
}

func (bar *bankAccountRepository) GetBalnceByIBANAndBIC(iban, bic string) (models.BankAccount, error) {
	var account models.BankAccount
	if err := bar.Where("iban = ? AND bic = ?", iban, bic).First(&account).Error; err != nil {
		return models.EmptyBankAccount, err
	}
	return account, nil
}

func (bar *bankAccountRepository) Update(account models.BankAccount) error {
	return bar.Save(account).Error
}
