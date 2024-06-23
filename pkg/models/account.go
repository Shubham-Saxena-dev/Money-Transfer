package models

import "fmt"

var (
	EmptyBankAccount = BankAccount{}
)

type BankAccount struct {
	ID               uint `gorm:"primaryKey"`
	OrganizationName string
	BalanceCents     int
	IBAN             string
	BIC              string
}

type Transfer struct {
	ID               uint `gorm:"primaryKey"`
	CounterpartyName string
	CounterpartyIBAN string
	CounterpartyBIC  string
	AmountCents      int
	BankAccountID    uint
	Description      string
}

func (b *BankAccount) HasSufficientBalance(deductibleAmountInCents int) bool {
	fmt.Println(b.BalanceCents, deductibleAmountInCents)
	return b.BalanceCents > deductibleAmountInCents
}

func (_ *BankAccount) TableName() string {
	return "bank_accounts"
}

func (_ *Transfer) TableName() string {
	return "transfers"
}
