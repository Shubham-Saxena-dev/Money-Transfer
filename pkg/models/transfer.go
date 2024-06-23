package models

import (
	log "github.com/sirupsen/logrus"
	"qonto/internal/customerrors"
	"strconv"
)

// TransferRequest represents a request for a transfer.
type TransferRequest struct {
	OrganizationIBAN string           `json:"organization_iban"`
	OrganizationBIC  string           `json:"organization_bic"`
	CreditTransfers  []CreditTransfer `json:"credit_transfers"`
}

type CreditTransfer struct {
	CounterpartyName string `json:"counterparty_name"`
	CounterpartyIBAN string `json:"counterparty_iban"`
	CounterpartyBIC  string `json:"counterparty_bic"`
	Amount           string `json:"amount"`
	Description      string `json:"description"`
}

func (t *TransferRequest) Validate() error {
	if len(t.OrganizationBIC) == 0 {
		return customerrors.ErrOrgBIC("organization_bic")
	}

	if len(t.OrganizationIBAN) == 0 {
		return customerrors.ErrOrgIBAN("organization_iban")
	}

	if t.CreditTransfers == nil || len(t.CreditTransfers) == 0 {
		return customerrors.ErrTransfer("credit_transfers")
	}

	for _, t := range t.CreditTransfers {
		if err := t.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (t *CreditTransfer) validate() error {
	if t.ConvertToCents() <= 0 {
		return customerrors.ErrValue("amount")
	}

	if len(t.CounterpartyBIC) == 0 {
		return customerrors.ErrOrgBIC("organization_bic")
	}

	if len(t.CounterpartyIBAN) == 0 {
		return customerrors.ErrOrgIBAN("organization_iban")
	}
	return nil
}

func (t *CreditTransfer) ConvertToCents() int {
	amount, err := strconv.ParseFloat(t.Amount, 64)
	if err != nil {
		log.Error(err)
		return 0
	}
	return int(amount * 100)
}
