package service

import (
	"qonto/internal/customerrors"
	"qonto/internal/repositories"
	"qonto/pkg/models"
)

type TransferService interface {
	ProcessTransfer(request models.TransferRequest) error
}

type transferService struct {
	repositories.TransactionRepository
}

func NewTransferService(transRepo repositories.TransactionRepository) TransferService {
	return &transferService{
		transRepo,
	}
}

func (t *transferService) ProcessTransfer(request models.TransferRequest) error {
	account, err := t.GetAccount(request.OrganizationIBAN, request.OrganizationBIC)
	if err != nil {
		return err
	}

	totalAmount := 0
	var dbTransfers []models.Transfer
	for _, transfer := range request.CreditTransfers {
		amount := transfer.ConvertToCents()

		dbTransfers = append(dbTransfers, models.Transfer{
			CounterpartyName: transfer.CounterpartyName,
			CounterpartyIBAN: transfer.CounterpartyIBAN,
			CounterpartyBIC:  transfer.CounterpartyBIC,
			AmountCents:      amount,
			BankAccountID:    account.ID,
			Description:      transfer.Description,
		})

		totalAmount += amount
	}

	if !account.HasSufficientBalance(totalAmount) {
		return customerrors.ErrBalance()
	}
	return t.DoTransfer(totalAmount, account, dbTransfers)
}
