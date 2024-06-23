package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"qonto/internal/customerrors"
	"qonto/internal/database"
	"qonto/pkg/models"
)

type TransactionRepository interface {
	GetAccount(string, string) (models.BankAccount, error)
	DoTransfer(int, models.BankAccount, []models.Transfer) error
}

type transactionRepository struct {
	*gorm.DB
	bankAccountRepo BankAccountRepository
}

func NewTransactionRepository(bankAccountRepo BankAccountRepository) TransactionRepository {
	return &transactionRepository{database.GetInstance(),
		bankAccountRepo}
}

func (t *transactionRepository) GetAccount(iban, bic string) (models.BankAccount, error) {
	account, err := t.bankAccountRepo.GetBalnceByIBANAndBIC(iban, bic)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.EmptyBankAccount, customerrors.ErrAcc(err)
	}
	return account, nil
}

func (t *transactionRepository) DoTransfer(totalAmt int, account models.BankAccount, transfers []models.Transfer) error {
	tx := t.Begin(&sql.TxOptions{Isolation: sql.LevelReadUncommitted})

	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("Panic occurred when transferring :%v", r))
			tx.Rollback()
		}
	}()

	account.BalanceCents -= totalAmt
	if err := tx.Save(account).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, transfer := range transfers {
		if err := tx.Create(&transfer).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
